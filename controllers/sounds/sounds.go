//go:generate go run ./generate_supported_sounds.go
package sounds

import (
	"fmt"
	"time"

	"github.com/circa10a/witchonstephendrive.com/internal/config"
	log "github.com/sirupsen/logrus"
)

const soundFileSuffix = ".mp3"
const homeAssistantContentType = "audio/mp3"
const homeAssistantMediaContextPath = "/local"
const homeAssistantCastContextPath = "/api/services/media_player/play_media"
const queuePollIntervalMS = 100
const soundPlaybackStatusPollIntervalSeconds = 3

// PlaySoundPayload is the type supported by home assistant to cast custom media
type PlaySoundPayload struct {
	EntityID         string `json:"entity_id"`
	MediaContentType string `json:"media_content_type"`
	MediaContentID   string `json:"media_content_id"`
}

// HomeAssistantStateResponse represents the /api/state/<entity_id> JSON response from home assistant
type HomeAssistantStateResponse struct {
	EntityID   string `json:"entity_id"`
	State      string `json:"state"`
	Attributes struct {
		VolumeLevel            float64     `json:"volume_level"`
		IsVolumeMuted          bool        `json:"is_volume_muted"`
		MediaContentID         string      `json:"media_content_id"`
		MediaDuration          float64     `json:"media_duration"`
		MediaPosition          float64     `json:"media_position"`
		MediaPositionUpdatedAt time.Time   `json:"media_position_updated_at"`
		AppID                  string      `json:"app_id"`
		AppName                string      `json:"app_name"`
		EntityPictureLocal     interface{} `json:"entity_picture_local"`
		FriendlyName           string      `json:"friendly_name"`
		SupportedFeatures      int         `json:"supported_features"`
	} `json:"attributes"`
	LastChanged time.Time `json:"last_changed"`
	LastUpdated time.Time `json:"last_updated"`
	Context     struct {
		ID       string      `json:"id"`
		ParentID interface{} `json:"parent_id"`
		UserID   string      `json:"user_id"`
	} `json:"context"`
}

// worker actually calls home assistant to play the sound read from the queue
func worker(witchConfig *config.WitchConfig, sound string) {
	log.Debug(fmt.Sprintf("playing sound: %s", sound))
	// Play sound
	_, err := witchConfig.HomeAssistantClient.R().SetBody(PlaySoundPayload{
		EntityID:         witchConfig.HomeAssistantEntityID,
		MediaContentType: homeAssistantContentType,
		MediaContentID:   fmt.Sprintf("%s/%s%s", homeAssistantMediaContextPath, sound, soundFileSuffix),
	}).Post(homeAssistantCastContextPath)

	// If enabled, check playback status until not "playing"
	if witchConfig.SoundQueueWaitUntilFinished {
		// Poll to determine when sound is finished
		for {
			time.Sleep(time.Second * soundPlaybackStatusPollIntervalSeconds)
			resp, err := witchConfig.HomeAssistantClient.R().
				SetResult(&HomeAssistantStateResponse{}).
				Get(fmt.Sprintf("api/states/%s", witchConfig.HomeAssistantEntityID))
			if err != nil {
				log.Error(err)
			}
			state := resp.Result().(*HomeAssistantStateResponse)
			log.Debug(fmt.Sprintf("waiting for entity id: %s to finish playing", state.EntityID))
			if state.State != "playing" {
				log.Debug(fmt.Sprintf("entity id: %s finished playing", state.EntityID))
				break
			}
		}
		if err != nil {
			log.Error(err)
		}
	}
}

// InitDaemon continually reads sounds out of a queue to ensure non-overlapping casting
func InitDaemon(witchConfig *config.WitchConfig) {
	for {
		// Don't waste cpu by polling the queue too much
		time.Sleep(time.Millisecond * queuePollIntervalMS)
		// Loop over queue, remove item, play, repeat
		var sounds = make([]string, witchConfig.SoundQueue.Size())
		log.Debug(fmt.Sprintf("sound queue size: %d", witchConfig.SoundQueue.Size()))
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
