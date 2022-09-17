package config

import (
	"github.com/circa10a/go-geofence"
	log "github.com/sirupsen/logrus"
)

// initGeofencingClient configures the geofencing client
func (w *WitchConfig) initGeofencingClient() {
	if w.GeofencingEnabled {
		client, err := geofence.New(&geofence.Config{
			IPAddress:               "",
			Token:                   w.GeofencingIPBaseAPIToken,
			Radius:                  w.GeofencingClient.Radius,
			AllowPrivateIPAddresses: true,
			CacheTTL:                -1, // hold cache until restart
		})
		if err != nil {
			log.Fatal(err)
		}
		w.GeofencingClient = *client
	}
}
