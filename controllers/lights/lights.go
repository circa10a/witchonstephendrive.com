package lights

import (
	"fmt"

	"github.com/amimof/huego"
	log "github.com/sirupsen/logrus"
)

const (
	hueManufacturerName  string = "Signify Netherlands B.V."
	innrManufacturerName string = "innr"
)

// SetLightsOn turns on all configured lights
func SetLightsOn(l []huego.Light) error {
	for _, light := range l {
		log.Debug(fmt.Sprintf("turning on light id: %d", light.ID))
		err := light.On()
		if err != nil {
			return err
		}
	}
	return nil
}

// SetLightsOff turns off all configured lights
func SetLightsOff(l []huego.Light) error {
	for _, light := range l {
		log.Debug(fmt.Sprintf("turning off light id: %d", light.ID))
		// Not all 3rd party manufacturers support setting power on behavior, innr supports this internally
		// This results in light colors being reset when being flashed if they don't support storing state
		if light.ManufacturerName == hueManufacturerName || light.ManufacturerName == innrManufacturerName {
			err := light.Off()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
