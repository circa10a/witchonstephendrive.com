package config

import (
	"fmt"

	"github.com/circa10a/witchonstephendrive.com/controllers/colors"
	"github.com/circa10a/witchonstephendrive.com/controllers/lights"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

// initHueLightsScheduler conditionally starts a scheduler to turn on/off lights and set default colors
func (w *WitchConfig) initHueLightsScheduler() {
	if w.HueLightsScheduleEnabled {
		// On
		log.Infof("scheduling lights turn on at hour: %d", w.HueLightsStart)
		log.Infof("scheduling lights turn off at hour: %d", w.HueLightsEnd)
		onSchedule := fmt.Sprintf("0 %d * * *", w.HueLightsStart)
		offschedule := fmt.Sprintf("0 %d * * *", w.HueLightsEnd)
		c := cron.New()
		// On
		_, err := c.AddFunc(onSchedule, func() {
			// If default colors are enabled and provided, turn on and set colors
			if w.HueDefaultColorsEnabled && len(w.HueDefaultColors) > 0 {
				log.Info("turning lights on and setting to default colors")
				err := colors.SetDefaultLightColors(w.HueDefaultColors, w.HueBridge)
				if err != nil {
					log.Error(err)
				}
			} else {
				// If just lights schedule is enabled, turn them on
				log.Info("turning lights on")
				err := lights.SetLightsOn(w.HueLightsStructs)
				if err != nil {
					log.Error(err)
				}
			}
		})
		if err != nil {
			log.Error(err)
		}
		// Off
		_, err = c.AddFunc(offschedule, func() {
			log.Info("turning lights off")
			for _, light := range w.HueLightsStructs {
				err := light.Off()
				if err != nil {
					log.Error(err)
				}
			}
		})
		if err != nil {
			log.Error(err)
		}
		c.Start()
	}
}
