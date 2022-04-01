//go:generate go run ./generate_supported_sounds.go
package sounds

import (
	"context"
	"fmt"
	"time"

	"github.com/circa10a/witchonstephendrive.com/internal/config"
	log "github.com/sirupsen/logrus"
)

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
	LastChanged time.Time `json:"last_changed"`
	LastUpdated time.Time `json:"last_updated"`
	Context     struct {
		ID       string      `json:"id"`
		ParentID interface{} `json:"parent_id"`
		UserID   string      `json:"user_id"`
	} `json:"context"`
	EntityID   string `json:"entity_id"`
	State      string `json:"state"`
	Attributes struct {
		MediaPositionUpdatedAt time.Time   `json:"media_position_updated_at"`
		EntityPictureLocal     interface{} `json:"entity_picture_local"`
		AppID                  string      `json:"app_id"`
		AppName                string      `json:"app_name"`
		FriendlyName           string      `json:"friendly_name"`
		MediaContentID         string      `json:"media_content_id"`
		MediaPosition          float64     `json:"media_position"`
		VolumeLevel            float64     `json:"volume_level"`
		MediaDuration          float64     `json:"media_duration"`
		SupportedFeatures      int         `json:"supported_features"`
		IsVolumeMuted          bool        `json:"is_volume_muted"`
	} `json:"attributes"`
}

// worker actually calls home assistant to play the sound read from the queue
func worker(witchConfig *config.WitchConfig, sound string) {
	log.Debug(fmt.Sprintf("Playing sound: %s", sound))
	// Play sound
	_, err := witchConfig.HomeAssistantClient.R().SetBody(PlaySoundPayload{
		EntityID:         witchConfig.HomeAssistantEntityID,
		MediaContentType: "audio/mp3",
		MediaContentID:   fmt.Sprintf("/local/%s.mp3", sound),
	}).Post("/api/services/media_player/play_media")
	if err != nil {
		// Attempts to play failed
		// Return so we don't unnecessarily try to get status of sound playing
		return
	}

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
		log.Debug(fmt.Sprintf("Waiting for entity id: %s to finish playing", state.EntityID))
		if state.State != "playing" {
			log.Debug(fmt.Sprintf("Entity id: %s finished playing", state.EntityID))
			break
		}
	}
	if err != nil {
		log.Error(err)
	}
}

// InitDaemon continually reads sounds out of a queue to ensure non-overlapping casting
func InitDaemon(ctx context.Context, w *config.WitchConfig) {
	ticker := time.NewTicker(queuePollIntervalMS)

	// Start daemon
	for {
		select {
		case <-ticker.C:
			log.Debugf("Queue size: %d", len(w.SoundQueue))
			select {
			case sound := <-w.SoundQueue:
				worker(w, sound)
			default:
				break
			}
		case <-ctx.Done():
			log.Info("Sound queue worker shutdown successfully")
			return
		}
	}
}

// IsDuringQuietTime ensures no sounds are played during configured/late hours
func IsDuringQuietTime(currentHour, startHour, endHour int) bool {
	return currentHour >= startHour || currentHour < endHour
}
