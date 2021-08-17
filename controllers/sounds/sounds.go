//go:generate go run ./generate_supported_sounds.go
package sounds

import (
	"fmt"
	"time"

	"github.com/circa10a/witchonstephendrive.com/internal/config"
	log "github.com/sirupsen/logrus"
)

const soundFileSuffix = ".mp3"
const assistantRelayCastType = "custom"
const assistantRelayCastContextPath = "/cast"
const queuePollIntervalMS = 100

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
		Source: fmt.Sprintf("%s%s", sound, soundFileSuffix),
		Type:   assistantRelayCastType,
	}).Post(assistantRelayCastContextPath)

	if err != nil {
		log.Error(err)
	}
}

// InitDaemon continually reads sounds out of a queue to ensure non-overlapping casting
func InitDaemon(witchConfig *config.WitchConfig) {
	for {
		// Don't waste cpu by polling the queue too much
		time.Sleep(time.Millisecond * queuePollIntervalMS)
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
