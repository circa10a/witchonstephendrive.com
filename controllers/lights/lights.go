package lights

import (
	"fmt"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	log "github.com/sirupsen/logrus"
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
func SetLightsOff(l []huego.Light, thirdPartyManufacturers []string) error {
	for _, light := range l {
		log.Debug(fmt.Sprintf("turning off light id: %d", light.ID))
		// Not all 3rd party manufacturers support setting power on behavior
		// This results in light colors being reset when being flashed
		if !utils.StrInSlice(light.ManufacturerName, thirdPartyManufacturers) {
			err := light.Off()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
