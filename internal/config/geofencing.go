package config

import (
	"github.com/qioalice/ipstack"
	log "github.com/sirupsen/logrus"
)

func (w *WitchConfig) InitGeofencing() {
	if w.GeofencingEnabled {
		if err := ipstack.Init(w.GeofencingIPStackAPIToken); err != nil {
			log.Error("Issue loading ip stack API token")
			log.Fatal(err)
		}
		// Get location of server
		currentLocation, err := ipstack.Me()
		if err != nil {
			log.Fatal(err)
		}
		// Cache location of server
		w.GeofencingCoordinates = currentLocation
		// Initialize map
		w.GeofencingCache = map[string]bool{}
	}
}
