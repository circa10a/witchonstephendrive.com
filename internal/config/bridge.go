package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/amimof/huego"
)

// InitHue discovers hue bridge and configured lights
func (w *WitchConfig) InitHueBridge() {
	w.Lock()
	defer w.Unlock()

	log.Info("Discovering bridge configuration")
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
}
