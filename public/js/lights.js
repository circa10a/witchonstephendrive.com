const setLight = async (color) => {
    try {
        await fetch(`/${color}`, {
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