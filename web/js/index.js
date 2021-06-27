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
    teal: "werewolf",
    rainbow: "this-is-halloween",
    default: "this-is-halloween"
};

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

// SetState (Main)
const setState = async (opts = {}) => {
    // Change navbar color
    let navBar =  document.getElementById("navbar");
    navBar.classList.remove(...navBar.classList);
    navBar.classList.add('nav-wrapper');
    navBar.classList.add(opts.color);
    sound = colorToSoundMap[opts.color] || colorToSoundMap['default'];

    // Set address bar theme color if supported
    document.querySelector('meta[name="theme-color"]').setAttribute('content', opts.color);

    // Play sound, don't wait since it takes a second to kick off
    try {
        fetch(`/sound/${sound}`, {
            method: 'POST',
        });
    } catch(e) {
        console.error(e);
    }

    // Set light colors via hue
    try {
        await fetch(`/color/${opts.color}`, {
            method: 'POST',
        });
    } catch(e) {
        console.error(e);
    }

    // Set footer text to include usage numbers
    let footer =  document.getElementById("footer");
    try {
        usageCount = await getCountOfColorChanges();
        footer.innerText = `Used ${usageCount} times`;
    } catch(e) {
        console.error(e);
    }
};
