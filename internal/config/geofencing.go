package config

import (
	"github.com/circa10a/go-geofence"
	log "github.com/sirupsen/logrus"
)

func (w *WitchConfig) InitGeofencing() {
	if w.GeofencingEnabled {
		client, err := geofence.New(&geofence.Config{
			IPAddress: "",
			Token:     w.GeofencingFreeGeoIPAPIToken,
			Radius:    w.GeofencingClient.Radius,
			CacheTTL:  -1, // hold cache until restart
		})
		if err != nil {
			log.Fatal(err)
		}
		w.GeofencingClient = *client
	}
}
