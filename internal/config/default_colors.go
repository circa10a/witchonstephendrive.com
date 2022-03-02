package config

import (
	"fmt"

	"github.com/circa10a/witchonstephendrive.com/controllers/colors"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

// initDefaultColorsScheduler conditionally starts a scheduler to set default colors on configured lights
func (w *WitchConfig) initDefaultColorsScheduler() {
	if w.HueDefaultColorsEnabled && len(w.HueDefaultColors) > 0 {
		log.Infof("scheduling default light colors to be set at hour: %d", w.HueDefaultColorsStart)
		schedule := fmt.Sprintf("0 %d * * *", w.HueDefaultColorsStart)
		c := cron.New()
		_, err := c.AddFunc(schedule, func() {
			log.Info("setting default colors")
			err := colors.SetDefaultLightColors(w.HueDefaultColors, w.HueBridge)
			if err != nil {
				log.Error(err)
			}
		})
		if err != nil {
			log.Error(err)
		}
		c.Start()
	}
}
