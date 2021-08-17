package lights

import "github.com/amimof/huego"

// SetLightsOn turns on all configured lights
func SetLightsOn(l []huego.Light) error {
	for _, light := range l {
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
		err := light.Off()
		if err != nil {
			return err
		}
	}
	return nil
}
