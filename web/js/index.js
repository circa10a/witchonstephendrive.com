// Base Path
const witchAPIBaseURL = '/api/v1';

// Utils
const sleep = (ms) => {
  return new Promise(resolve => setTimeout(resolve, ms));
};

const checkLimit = (opts = {}) => {
  if (opts.statusCode === 429) {
    M.toast({html: 'Slow down!', displayLength: 2000, classes: 'green darken-1 rounded'});
  };
};

// Sounds
const playRandomSound = async() => {
  // Get Supported Sounds
  // Have a default
  supportedSounds = ['this-is-halloween'];
  try {
    let resp = await fetch(`${witchAPIBaseURL}/sounds`, {});
    let json = await resp.json();
    supportedSounds = json.supportedSounds;
  } catch(e) {
    console.error(e);
  }

  // Let's not let users play this one
  let allowedSounds = supportedSounds.filter(sound => sound !== 'police-siren');
  const randomSound = () => allowedSounds[Math.floor(Math.random() * allowedSounds.length)];

  try {
    let resp = await fetch(`${witchAPIBaseURL}/sound/${randomSound()}`, {
      method: 'POST',
    });
    let json = await resp.json();
    if (json.message.includes('quiet')) {
      M.toast({html: json.message, displayLength: 2000, classes: 'green darken-1 rounded'});
    }
  } catch(e) {
    console.error(e);
  }
};

// Metrics
const getCountOfColorChanges = async () => {
  metricsResp = await fetch('/metrics', {
    method: 'GET',
  });
  const metricsText = await metricsResp.text();
  const parsed = parsePrometheusTextFormat(metricsText);
  const requestMetrics = parsed.filter(metric =>  metric.name ===  'echo_requests_total')[0];
  const requestCounters = requestMetrics.metrics.filter(counter => {
    return counter.labels.code === '200' &&
           counter.labels.method === 'POST' &&
           counter.labels.url.includes('/color/') &&
           counter.labels.host === 'witchonstephendrive.com';
  });
  let usageCount = 0;
  for (counter of requestCounters) {
    usageCount += Number(counter.value);
  }
  return usageCount;
};

const setUsageFooter = async () => {
  // Set footer text to include usage numbers
  let footer =  document.getElementById('footer');
  try {
    usageCount = await getCountOfColorChanges();
    footer.innerText = `Used ${usageCount} times`;
  } catch(e) {
    console.error(e);
  }
};

// Theming UI colors
const setTheme = (opts = {}) => {
  // Change navbar color
  let navBar =  document.getElementById('navbar');
  navBar.classList.remove(...navBar.classList);
  navBar.classList.add('nav-wrapper');
  navBar.classList.add(opts.color);

  // Set address bar theme color if supported
  document.querySelector('meta[name="theme-color"]').setAttribute('content', opts.color);
};

// LightState
const setState = async (opts = {}) => {
  // Set UI colors
  setTheme(opts);
  // Play sound, don't wait since it takes a second to kick off
  playRandomSound();
  // Set light colors via hue
  try {
    colorResponse = await fetch(`${witchAPIBaseURL}/color/${opts.color}`, {
      method: 'POST',
    });
    // Check that rate limit isn't hit, alert user
    checkLimit({statusCode: colorResponse.status});
  } catch(e) {
    console.error(e);
  }
  setUsageFooter();
};

// Flicker (turn lights off/on)
const flicker = async (opts = {count: 3, sleepTime: 1000, color: 'black'}) => {
  setTheme({color: opts.color});
  // Play sound, don't wait since it takes a few seconds to kick off
  playRandomSound();
  // Try to sync up sound and flash
  await sleep(2000);
  for (let i = 0; i < opts.count; i++) {
    // Play sound, don't wait since it takes a second to kick off
    try {
      await fetch(`${witchAPIBaseURL}/lights/off`, {
        method: 'POST',
      });
      await sleep(opts.sleepTime);
      await fetch(`${witchAPIBaseURL}/lights/on`, {
        method: 'POST',
      });
      await sleep(opts.sleepTime);
    } catch(e) {
      console.error(e);
    };
  };
  setUsageFooter();
};
