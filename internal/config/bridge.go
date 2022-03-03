package config

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/amimof/huego"
)

// initHue discovers hue bridge and configured lights on configured interval
func (w *WitchConfig) initHue() {
	// Run this regularly in the event the bridge gets a new IP address
	for {
		w.initHueBridge()
		time.Sleep(w.HueBridgeRefreshInterval)
	}
}

func (w *WitchConfig) initHueBridge() {
	w.mu.Lock()
	defer w.mu.Unlock()

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
