package config

import (
	"fmt"

	"github.com/circa10a/witchonstephendrive.com/controllers/lights"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func (w *WitchConfig) InitHueLightsScheduler() {
	if w.HueLightsScheduleEnabled {
		// On
		log.Infof("scheduling lights turn on at hour: %d", w.HueLightsStart)
		log.Infof("scheduling lights turn off at hour: %d", w.HueLightsEnd)
		onSchedule := fmt.Sprintf("0 %d * * *", w.HueLightsStart)
		offschedule := fmt.Sprintf("0 %d * * *", w.HueLightsEnd)
		c := cron.New()
		// On
		_, err := c.AddFunc(onSchedule, func() {
			log.Info("turning lights on")
			err := lights.SetLightsOn(w.HueLightsStructs)
			if err != nil {
				log.Error(err)
			}
		})
		if err != nil {
			log.Error(err)
		}
		// Off
		_, err = c.AddFunc(offschedule, func() {
			log.Info("turning lights off")
			err := lights.SetLightsOff(w.HueLightsStructs, w.ThirdPartyManufacturers)
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
