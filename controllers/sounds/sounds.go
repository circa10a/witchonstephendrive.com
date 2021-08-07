package sounds

import (
	"fmt"
	"time"

	"github.com/circa10a/witchonstephendrive.com/internal/config"
	log "github.com/sirupsen/logrus"
)

// SupportedSounds is a slice of available mp3's in the sounds directory
var SupportedSounds = []string{
	"adams-family",
	"dracula",
	"ghost",
	"halloween-organ",
	"leave-now",
	"police-siren", // Cause Traci
	"pumpkin-king",
	"scream",
	"spell-on-you",
	"stranger-things",
	"this-is-halloween",
	"werewolf",
	"witch-laugh",
	"youll-float-too",
}

// PlaySoundPayload is the type supported by assistant-relay to cast custom media
type PlaySoundPayload struct {
	Device string `json:"device"`
	Source string `json:"source"`
	Type   string `json:"type"`
}

// worker actually calls the assistant relay to play the sound read from the queue
func worker(witchConfig *config.WitchConfig, sound string) {
	_, err := witchConfig.RelayClient.R().SetBody(PlaySoundPayload{
		Device: witchConfig.AssistantDevice,
		Source: fmt.Sprintf("%v.mp3", sound),
		Type:   "custom",
	}).Post("/cast")

	if err != nil {
		log.Error(err)
	}
}

// Daemon continually reads sounds out of a queue to ensure non-overlapping casting
func Daemon(witchConfig *config.WitchConfig) {
	for {
		// Don't waste cpu by polling the queue too much
		time.Sleep(time.Second * time.Duration(witchConfig.SoundQueuePollInterval))
		// Loop over queue, remove item, play, repeat
		var sounds = make([]string, witchConfig.SoundQueue.Size())
		for i := 0; i < len(sounds); i++ {
			sounds[i] = witchConfig.SoundQueue.Shift().(string)
			worker(witchConfig, sounds[i])
		}
	}
}

// IsDuringQuietTime ensures no sounds are played during configured/late hours
func IsDuringQuietTime(currentHour, startHour, endHour int) bool {
	return currentHour >= startHour || currentHour < endHour
}
