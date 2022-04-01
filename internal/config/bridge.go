package config

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/amimof/huego"
)

// InitHue discovers hue bridge and configured lights on configured interval
func (w *WitchConfig) InitHue(ctx context.Context) {
	// Run this regularly in the event the bridge gets a new IP address
	ticker := time.NewTicker(w.HueBridgeRefreshInterval)
	// Set initial address
	w.initHueBridge()
	// Start daemon
	for {
		select {
		case <-ticker.C:
			w.initHueBridge()
		case <-ctx.Done():
			log.Info("Hue bridge address refresher shutdown successfully")
			return
		}
	}
}

func (w *WitchConfig) initHueBridge() {
	w.mu.Lock()
	defer w.mu.Unlock()

	log.Info("Renewing bridge configuration")
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
