const setLight = async (color) => {
    const rawResponse = await fetch(`https://maker.ifttt.com/trigger/hue-hook/with/key/lMsWNtWo61JsG5pXq2FCrmlgdX1o_9eRDukbNFgQLQk`, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ 'value1': color }),
    });
    const content = await rawResponse.json();

    console.log(content);
};

const lightDance = async (colors) => {
    count = 0;
    threshold = 5;
    while (count < threshold) {
        for (color of colors) {
            await setLight(color);
            await sleep(6000);
            count++;
        };
        console.log(count)
    };
};