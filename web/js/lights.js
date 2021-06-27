const setState = async (opts = {}) => {
    // Change navbar
    let navBar =  document.getElementById("navbar");
    navBar.classList.remove(...navBar.classList);
    navBar.classList.add('nav-wrapper');
    navBar.classList.add(opts.color);
    sound = colorToSoundMap[opts.color] ?? colorToSoundMap['default'];
    // Play sound, don't wait since it takes a second to kick off
    try {
        fetch(`/sound/${sound}`, {
            method: 'POST',
        });
    } catch(e) {
        console.error(e)
    }

    // Set light colors via hue
    try {
        await fetch(`/color/${opts.color}`, {
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