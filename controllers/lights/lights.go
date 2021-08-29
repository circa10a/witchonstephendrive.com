package lights

import (
	"fmt"

	"github.com/amimof/huego"
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
func SetLightsOff(l []huego.Light) error {
	for _, light := range l {
		log.Debug(fmt.Sprintf("turning off light id: %d", light.ID))
		err := light.Off()
		if err != nil {
			return err
		}
	}
	return nil
}
