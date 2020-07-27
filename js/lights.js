const setLight = async (color) => {
    const rawResponse = await fetch(`http://${endpoint}/api/${user}/lights/${light}/state`, {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ 'on': true, 'bri': 200, "hue": parseInt(color) }),
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
            await sleep(3000);
            count++;
        };
        console.log(count)
    };
};