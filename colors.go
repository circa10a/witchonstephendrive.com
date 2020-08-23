package main

import "github.com/amimof/huego"

var colorMap = map[string]huego.State{
	"red": {
		On:  true,
		Bri: 254,
		Xy: []float32{
			0.6786,
			0.3126,
		},
	},
	"orange": {
		On:  true,
		Bri: 254,
		Xy: []float32{
			0.6424,
			0.3523,
		},
	},
	"yellow": {
		On:  true,
		Bri: 254,
		Xy: []float32{
			0.5145,
			0.4691,
		},
	},
	"green": {
		On:  true,
		Bri: 254,
		Xy: []float32{
			0.1723,
			0.6866,
		},
	},
	"blue": {
		On:  true,
		Bri: 254,
		Xy: []float32{
			0.1396,
			0.0686,
		},
	},
	"purple": {
		On:  true,
		Bri: 254,
		Xy: []float32{
			0.2027,
			0.0715,
		},
	},
	"pink": {
		On:  true,
		Bri: 254,
		Xy: []float32{
			0.3924,
			0.1634,
		},
	},
	"teal": {
		On:  true,
		Bri: 254,
		Xy: []float32{
			0.1527,
			0.2144,
		},
	},
}

func getColorState(c string) huego.State {
	return colorMap[c]
}