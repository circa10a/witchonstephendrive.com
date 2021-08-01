package sounds

import (
	"fmt"
	"net/http"

	"github.com/circa10a/witchonstephendrive.com/internal/config"
	log "github.com/sirupsen/logrus"
)

// SupportedSounds is a slice of available mp3's in the sounds directory
var SupportedSounds = []string{
	"dracula",
	"ghost",
	"halloween-organ",
	"leave-now",
	"pumpkin-king",
	"scream",
	"this-is-halloween",
	"werewolf",
	"witch-laugh",
	"police-siren", // Cause Traci
}

// PlaySoundPayload is the type supported by assistant-relay to cast custom media
type PlaySoundPayload struct {
	Device string `json:"device"`
	Source string `json:"source"`
	Type   string `json:"type"`
}

// Worker reads sounds out of a channel to ensure non-overlapping casting
func Worker(witchConfig config.WitchConfig) {
	for sound := range witchConfig.SoundChannel {
		resp, err := witchConfig.Client.R().SetBody(PlaySoundPayload{
			Device: witchConfig.AssistantDevice,
			Source: fmt.Sprintf("%v.mp3", sound),
			Type:   "custom",
		}).Post("/cast")

		if err != nil {
			log.Error(err)
		}
		if resp.StatusCode() != http.StatusOK {
			log.Error(resp.Body())
		}
	}
}
