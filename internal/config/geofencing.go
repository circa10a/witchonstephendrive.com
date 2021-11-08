package config

import (
	"github.com/circa10a/go-geofence"
	log "github.com/sirupsen/logrus"
)

func (w *WitchConfig) InitGeofencing() {
	if w.GeofencingEnabled {
		client, err := geofence.New("", w.GeofencingIPStackAPIToken, w.GeofencingSensitivity)
		if err != nil {
			log.Fatal(err)
		}
		w.GeofencingClient.CreateCache(-1)
		w.GeofencingClient = *client
	}
}
