package config

import (
	"fmt"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"

	"github.com/amimof/huego"
)

// InitHue discovers hue bridge and configured lights on configured interval
func (w *WitchConfig) InitHue() {
	// Run this regularly in the event the bridge gets a new IP address
	schedule := fmt.Sprintf("0 */%d * * * ", w.HueBridgeRefreshInterval)
	c := cron.New()
	_, err := c.AddFunc(schedule, func() {
		w.initHueBridge()
	})
	if err != nil {
		log.Error(err)
	}
	w.initHueBridge()
	c.Start()
}

func (w *WitchConfig) initHueBridge() {
	log.Info("renewing bridge configuration")
	// Find hue bridge ip
	hueBridge, err := huego.Discover()
	if err != nil {
		log.Fatal(err)
	}
	w.HueBridge = hueBridge
	// Authenticate against bridge api
	w.HueBridge.Login(w.HueToken)
	// Store all light data to be used later
	for _, lightID := range w.HueLights {
		light, err := w.HueBridge.GetLight(lightID)
		if err != nil {
			log.Error(err)
		}
		w.HueLightsStructs = append(w.HueLightsStructs, *light)
	}
	if err != nil {
		log.Error(err)
	}
}
