package colors

import (
	"errors"
	"fmt"
	"sort"

	"github.com/amimof/huego"
	log "github.com/sirupsen/logrus"
)

// ColorMap holds all the possible colors supported by the api
type ColorMap map[string]huego.State

var maxBrightness uint8 = 254
var colorMode string = "xy"
var defaultEffect string = "none"

// SupportedColors hold a list of supported colors
var SupportedColors = getSupportedColors()

// Colors hold all the supported colors' states
var Colors = ColorMap{
	"blue": {
		On:        true,
		Bri:       maxBrightness,
		ColorMode: colorMode,
		Effect:    defaultEffect,
		Xy: []float32{
			0.1396,
			0.0686,
		},
	},
	"green": {
		On:        true,
		Bri:       maxBrightness,
		ColorMode: colorMode,
		Effect:    defaultEffect,
		Xy: []float32{
			0.1723,
			0.6866,
		},
	},
	"orange": {
		On:        true,
		Bri:       maxBrightness,
		ColorMode: colorMode,
		Effect:    defaultEffect,
		Xy: []float32{
			0.6424,
			0.3523,
		},
	},
	"pink": {
		On:        true,
		Bri:       maxBrightness,
		ColorMode: colorMode,
		Effect:    defaultEffect,
		Xy: []float32{
			0.3924,
			0.1634,
		},
	},
	"purple": {
		On:        true,
		Bri:       maxBrightness,
		ColorMode: colorMode,
		Effect:    defaultEffect,
		Xy: []float32{
			0.2027,
			0.0715,
		},
	},
	"rainbow": {
		On:        true,
		Bri:       maxBrightness,
		ColorMode: colorMode,
		Effect:    "colorloop",
	},
	"red": {
		On:        true,
		Bri:       maxBrightness,
		ColorMode: colorMode,
		Effect:    defaultEffect,
		Xy: []float32{
			0.6786,
			0.3126,
		},
	},
	"teal": {
		On:        true,
		Bri:       maxBrightness,
		ColorMode: colorMode,
		Effect:    defaultEffect,
		Xy: []float32{
			0.1527,
			0.2144,
		},
	},
	"yellow": {
		On:        true,
		Bri:       maxBrightness,
		ColorMode: colorMode,
		Effect:    defaultEffect,
		Xy: []float32{
			0.5145,
			0.4691,
		},
	},
}

// getSupportedColors returns a slice of supported colors
func getSupportedColors() []string {
	supportedColors := []string{}
	for color := range Colors {
		supportedColors = append(supportedColors, color)
	}
	sort.Strings(supportedColors)
	return supportedColors
}

// ErrColorNotSupported error gets raised in the event that a requested color is not in the Colors map
var ErrColorNotSupported = errors.New("color not supported")

// SetLightsColor sets all configured lights to the same color
func SetLightsColor(lights []int, bridge *huego.Bridge, color string) error {
	if _, ok := Colors[color]; ok {
		for _, light := range lights {
			log.Debug(fmt.Sprintf("setting color: %s on light id: %d", color, light))
			_, err := bridge.SetLightState(light, Colors[color])
			if err != nil {
				return err
			}
		}
	} else {
		return ErrColorNotSupported
	}
	return nil
}

// SetDefaultLightColors sets the default configured colors initiated by schedule
func SetDefaultLightColors(defaultColorsMap map[int]string, bridge *huego.Bridge) error {
	for light, color := range defaultColorsMap {
		log.Debug(fmt.Sprintf("setting default color: %s on light id: %d", color, light))
		if _, ok := Colors[color]; ok {
			_, err := bridge.SetLightState(light, Colors[color])
			if err != nil {
				return err
			}
		} else {
			return ErrColorNotSupported
		}
	}
	return nil
}
