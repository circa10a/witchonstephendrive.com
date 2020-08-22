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
    threshold = 10;
    while (count < threshold) {
        for (color of colors) {
            await setLight(color);
            await sleep(1000);
            count++;
        };
    };
};