const setColors = async (opts = {}) => {
    // Change navbar
    let navBar =  document.getElementById("navbar");
    navBar.classList.remove(...navBar.classList);
    navBar.classList.add('nav-wrapper');
    navBar.classList.add(opts.themeColor === 'rainbow' ? 'black' : opts.themeColor);
    // Set light colors via hue
    try {
        await fetch(`/color/${opts.lightsColor}`, {
            method: 'POST',
        });
    } catch(e) {
        console.error(e)
    }
};

const lightDance = async (colors) => {
    count = 0;
    while (count < colors.length) {
        for (color of colors) {
            await setLight(color);
            await sleep(1000);
            count++;
        };
    };
};