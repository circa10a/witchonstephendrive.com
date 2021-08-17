package sounds

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/circa10a/witchonstephendrive.com/internal/config"
	log "github.com/sirupsen/logrus"
)

const soundsDirectory = "./sounds"
const soundFileSuffix = ".mp3"
const assistantRelayCastType = "custom"
const assistantRelayCastContextPath = "/cast"

// SupportedSounds is a slice of available mp3's in the sounds directory
var SupportedSounds = getSupportedSounds()

// getSupportedColors returns a slice of supported sounds sourced from the sounds directory
func getSupportedSounds() []string {
	supportedSounds := []string{}
	files, err := ioutil.ReadDir(soundsDirectory)
	if err != nil {
		log.Error(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), soundFileSuffix) {
			// dracula.mp3 => dracula
			supportedSounds = append(supportedSounds, strings.TrimSuffix(file.Name(), soundFileSuffix))
		}
	}
	return supportedSounds
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
		Source: fmt.Sprintf("%s%s", sound, soundFileSuffix),
		Type:   assistantRelayCastType,
	}).Post(assistantRelayCastContextPath)

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
