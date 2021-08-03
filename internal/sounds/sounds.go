package sounds

import (
	"fmt"
	"time"

	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/nleeper/goment"
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
	"spell-on-you",
	"stranger-things",
	"adams-family",
	"police-siren", // Cause Traci
}

// PlaySoundPayload is the type supported by assistant-relay to cast custom media
type PlaySoundPayload struct {
	Device string `json:"device"`
	Source string `json:"source"`
	Type   string `json:"type"`
}

// worker actually calls the assistant relay to play the sound read from the queue
func worker(witchConfig config.WitchConfig, sound string) {
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
func Daemon(witchConfig config.WitchConfig) {
	for {
		time.Sleep(time.Second * time.Duration(witchConfig.SoundQueuePollInterval))
		// If there is something at the front of the queue, remove it, and play it
		if witchConfig.SoundQueue.Head() != nil {
			sound := witchConfig.SoundQueue.Dequeue()
			worker(witchConfig, sound.(string))
		}
	}
}

// IsDuringQuietTime ensures no sounds are played during configured/late hours
func IsDuringQuietTime(startTime, endTime string) bool {
	token := "LT"
	now, err := goment.New(time.Now())
	if err != nil {
		log.Error(err)
	}
	start, err := goment.New(startTime, token)
	if err != nil {
		log.Error(err)
	}
	end, err := goment.New(endTime, token)
	if err != nil {
		log.Error(err)
	}
	return now.IsBetween(start, end)
}
