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

const (
	maxBrightness uint8  = 254
	defaultEffect string = "none"
	// How many times we'll set the XY color to get to desired state if manufacturer not supported
	thirdPartyManufacturerRetryCount int    = 16
	hueManufacturerName              string = "Signify Netherlands B.V."
	innrManufacturerName             string = "innr"
)

// SupportedColors hold a list of supported colors
var SupportedColors = getSupportedColors()

// Colors hold all the supported colors' states
var Colors = ColorMap{
	"blue": {
		On:     true,
		Bri:    maxBrightness,
		Effect: defaultEffect,
		Xy: []float32{
			0.1396,
			0.0686,
		},
	},
	"cyan": {
		On:     true,
		Bri:    maxBrightness,
		Effect: defaultEffect,
		Xy: []float32{
			0.1527,
			0.2144,
		},
	},
	"green": {
		On:     true,
		Bri:    maxBrightness,
		Effect: defaultEffect,
		Xy: []float32{
			0.1723,
			0.6866,
		},
	},
	"orange": {
		On:     true,
		Bri:    maxBrightness,
		Effect: defaultEffect,
		Xy: []float32{
			0.6424,
			0.3523,
		},
	},
	"pink": {
		On:     true,
		Bri:    maxBrightness,
		Effect: defaultEffect,
		Xy: []float32{
			0.3924,
			0.1634,
		},
	},
	"purple": {
		On:     true,
		Bri:    maxBrightness,
		Effect: defaultEffect,
		Xy: []float32{
			0.2027,
			0.0715,
		},
	},
	"rainbow": {
		On:     true,
		Bri:    maxBrightness,
		Effect: "colorloop",
	},
	"red": {
		On:     true,
		Bri:    maxBrightness,
		Effect: defaultEffect,
		Xy: []float32{
			0.6786,
			0.3126,
		},
	},
	"yellow": {
		On:     true,
		Bri:    maxBrightness,
		Effect: defaultEffect,
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
func SetLightsColor(lights []huego.Light, bridge *huego.Bridge, color string) []error {
	if _, ok := Colors[color]; ok {
		errs := []error{}
		for _, light := range lights {
			state := Colors[color]
			log.Debug(fmt.Sprintf("Setting color: %s on light id: %d", color, light.ID))
			_, err := bridge.SetLightState(light.ID, state)
			if err != nil {
				errs = append(errs, err)
			}
			// Effect seems to clash with innr lights and doesn't set color properly
			if light.ManufacturerName == innrManufacturerName {
				state.Effect = ""
				_, err := bridge.SetLightState(light.ID, state)
				if err != nil {
					errs = append(errs, err)
				}
			}
			// Not all 3rd party manufacturers support setting the light state in one apply
			// But requires multiple calls to set state to change to the desired colors
			if light.ManufacturerName != hueManufacturerName && light.ManufacturerName != innrManufacturerName {
				for i := 0; i < thirdPartyManufacturerRetryCount; i++ {
					_, err := bridge.SetLightState(light.ID, state)
					if err != nil {
						errs = append(errs, err)
					}
				}
			}
		}
		if len(errs) > 0 {
			return errs
		}
	} else {
		return []error{ErrColorNotSupported}
	}
	return nil
}

// SetDefaultLightColors sets the default configured colors initiated by schedule
func SetDefaultLightColors(defaultColorsMap map[int]string, bridge *huego.Bridge) error {
	for light, color := range defaultColorsMap {
		log.Debug(fmt.Sprintf("Setting default color: %s on light id: %d", color, light))
		if _, ok := Colors[color]; ok {
			light, err := bridge.GetLight(light)
			if err != nil {
				return err
			}
			errs := SetLightsColor([]huego.Light{*light}, bridge, color)
			if len(errs) > 0 {
				return errs[0]
			}
		} else {
			return ErrColorNotSupported
		}
	}
	return nil
}
