// Base Path
const witchAPIBaseURL = '/api/v1';

// Utils
const sleep = (ms) => {
  return new Promise(resolve => setTimeout(resolve, ms));
};

const checkIfAllowed = (opts = {}) => {
  // Rate limiting
  if (opts.statusCode === 429) {
    M.toast({html: 'Slow down!', displayLength: 2000, classes: 'green darken-1 rounded'});
  };
  // Geofencing
  if (opts.statusCode === 403) {
    M.toast({html: 'It appears you are not on Stephen Drive. Come on over to have some fun!', displayLength: 3000, classes: 'green darken-1 rounded'});
  };
};

const convertWitchTimeToLocalTimeMessage = (apiResponse) => {
  const witchTZ = 'America/Chicago';
  const userTZ = dayjs.tz.guess();
  const outputDateFormat = 'h:mmA';
  const [startTime, endTime] = apiResponse.match(/\d+:00/g);
  const formattedDate = dayjs().format('YYYY-DD-MM');
  const localStartTime = dayjs.tz(`${formattedDate} ${startTime}`, witchTZ).tz(userTZ).format(outputDateFormat);
  const localEndTime = dayjs.tz(`${formattedDate} ${endTime}`, witchTZ).tz(userTZ).format(outputDateFormat);
  return `sounds disabled. quiet time is between ${localStartTime} and ${localEndTime}`;
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
      const userMsg = convertWitchTimeToLocalTimeMessage(json.message);
      M.toast({html: userMsg, displayLength: 2000, classes: 'green darken-1 rounded'});
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
  // Play sound, dont' wait
  playRandomSound();
  // Set light colors via hue
  try {
    colorResponse = await fetch(`${witchAPIBaseURL}/color/${opts.color}`, {
      method: 'POST',
    });
    // Check that rate limit isn't hit, alert user
    // Check if allowed via geofencing
    checkIfAllowed({statusCode: colorResponse.status});
  } catch(e) {
    console.error(e);
  }
  // setUsageFooter();
};

// Flicker (turn lights off/on)
const flicker = async (opts = {count: 3, sleepTime: 1000, color: 'black'}) => {
  setTheme({color: opts.color});
  // Play sound, don't wait
  playRandomSound();
  for (let i = 0; i < opts.count; i++) {
    // Play sound, don't wait
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
  // setUsageFooter();
};
