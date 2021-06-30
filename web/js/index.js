// Utils
function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

// Sounds
const colorToSoundMap = {
    red: "dracula",
    orange: "pumpkin-king",
    yellow: "scream",
    green: "witch-laugh",
    blue: "ghost",
    purple: "halloween-organ",
    pink: "leave-now",
    rainbow: "this-is-halloween",
    black: "werewolf",
    default: "this-is-halloween"
};

// Sounds
const playSoundForColor = async (opts = {}) => {
    sound = colorToSoundMap[opts.color] || colorToSoundMap['default'];
    try {
        await fetch(`/sound/${sound}`, {
            method: 'POST',
        });
    } catch(e) {
        console.error(e);
    }
}

// Metrics
const getCountOfColorChanges = async () => {
    metricsResp = await fetch('/metrics', {
        method: 'GET',
    });

    const metricsText = await metricsResp.text();
    const parsed = parsePrometheusTextFormat(metricsText);
    const requestMetrics = parsed.filter(metric =>  metric.name ===  'echo_requests_total')[0];
    const requestCounter = requestMetrics.metrics.filter(counter => {
        return counter.labels.code === "200" && counter.labels.url === '/color/:color' && counter.labels.host === 'witchonstephendrive.com';
    })[0];

    return requestCounter.value;
};

const setUsageFooter = async () => {
     // Set footer text to include usage numbers
     let footer =  document.getElementById("footer");
     try {
         usageCount = await getCountOfColorChanges();
         footer.innerText = `Used ${usageCount} times`;
     } catch(e) {
         console.error(e);
     }
}

// Theming UI colors
const setTheme = (opts = {}) => {
    sound = colorToSoundMap[opts.color] || colorToSoundMap['default'];

    // Change navbar color
    let navBar =  document.getElementById("navbar");
    navBar.classList.remove(...navBar.classList);
    navBar.classList.add('nav-wrapper');
    navBar.classList.add(opts.color);

    // Set address bar theme color if supported
    document.querySelector('meta[name="theme-color"]').setAttribute('content', opts.color);
}

// LightState
const setState = async (opts = {}) => {
    // Set UI colors
    setTheme(opts);
    // Play sound, don't wait since it takes a second to kick off
    playSoundForColor(opts)

    // Set light colors via hue
    try {
        await fetch(`/color/${opts.color}`, {
            method: 'POST',
        });
    } catch(e) {
        console.error(e);
    }
    setUsageFooter()
};

// Flicker (turn lights off/on)
const flicker = async (opts = {count: 2, sleepTime: 1000, color: 'black'}) => {
    setTheme({color: opts.color})
    // Play sound, don't wait since it takes a few seconds to kick off
    playSoundForColor(opts)
    // Try to sync up sound and flash
    await sleep(5000)
    for (let i = 0; i < opts.count; i++) {
        // Play sound, don't wait since it takes a second to kick off
        try {
            await fetch('/lights/off', {
                method: 'POST',
            });
            await sleep(opts.sleepTime);
            await fetch('/lights/on', {
                method: 'POST',
            });
            await sleep(opts.sleepTime);
        } catch(e) {
            console.error(e);
        }
      }
    setUsageFooter()
}
