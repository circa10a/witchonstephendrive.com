package colors

import "github.com/amimof/huego"

// ColorMap holds all the possible colors supported by the api
type ColorMap map[string]huego.State

var maxBrightness uint8 = 254

// SupportedColors hold a list of supported colors
var SupportedColors = GetSupportedColors()

// Colors hold all the supported colors' states
var Colors = ColorMap{
	"red": {
		On:     true,
		Bri:    maxBrightness,
		Effect: "none",
		Xy: []float32{
			0.6786,
			0.3126,
		},
	},
	"orange": {
		On:     true,
		Bri:    maxBrightness,
		Effect: "none",
		Xy: []float32{
			0.6424,
			0.3523,
		},
	},
	"yellow": {
		On:     true,
		Bri:    maxBrightness,
		Effect: "none",
		Xy: []float32{
			0.5145,
			0.4691,
		},
	},
	"green": {
		On:     true,
		Bri:    maxBrightness,
		Effect: "none",
		Xy: []float32{
			0.1723,
			0.6866,
		},
	},
	"blue": {
		On:     true,
		Bri:    maxBrightness,
		Effect: "none",
		Xy: []float32{
			0.1396,
			0.0686,
		},
	},
	"purple": {
		On:     true,
		Bri:    maxBrightness,
		Effect: "none",
		Xy: []float32{
			0.2027,
			0.0715,
		},
	},
	"pink": {
		On:     true,
		Bri:    maxBrightness,
		Effect: "none",
		Xy: []float32{
			0.3924,
			0.1634,
		},
	},
	"rainbow": {
		On:     true,
		Bri:    maxBrightness,
		Effect: "colorloop",
	},
}

// GetSupportedColors returns a slice of supported colors
func GetSupportedColors() []string {
	supportedColors := []string{}
	for color := range Colors {
		supportedColors = append(supportedColors, color)
	}
	return supportedColors
}
