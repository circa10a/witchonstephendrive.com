package routes

import (
	"fmt"
	"net/http"

	"github.com/circa10a/witchonstephendrive.com/internal/sounds"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// SoundsListResponse responds supported sounds to play
type SoundsListResponse struct {
	SupportedSounds []string `json:"supportedSounds"`
}

// SoundSuccesfulPlayResponse responds with a success status string when operation has completed successfully
type SoundSuccessfulPlayResponse struct {
	Status string `json:"status"`
}

// SoundFailedPlayResponse responds with status string, reason for failure in message, and list of supported sounds
type SoundFailedPlayResponse struct {
	Status          string   `json:"status"`
	Message         string   `json:"message"`
	SupportedSounds []string `json:"supportedSounds"`
}

// sounds godoc
// @Summary Get available sounds to play
// @Description Get list of supported sounds
// @Produce json
// @Success 200 {object} SoundsListResponse
// @Router /sounds [get]
func soundsReadHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, SoundsListResponse{
		SupportedSounds: sounds.SupportedSounds,
	})
}

// :sound godoc
// @Summary Play sound via assistant relay
// @Description Play halloween sound supported in sound list
// @Produce json
// @Success 200 {object} SoundSuccessfulPlayResponse
// @Failure 400 {object} SoundFailedPlayResponse
// @Failure 500 {object} SoundFailedPlayResponse
// @Router /sound/{sound} [post]
// @Param sound path string true "Sound to play"
func soundPlayHandler(c echo.Context) error {
	sound := c.Param("sound")
	client := c.Get("client").(*resty.Client)
	assistantDevice := c.Get("assistantDevice")
	if utils.StrInSlice(sound, sounds.SupportedSounds) {
		// Call assistant-relay if sound is supported
		resp, err := client.R().SetBody(&sounds.PlaySoundPayload{
			Device: assistantDevice.(string),
			Source: fmt.Sprintf("%v.mp3", sound),
			Type:   "custom",
		}).Post("/cast")
		// Handle unknown error
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, SoundFailedPlayResponse{
				Status:          failedString,
				Message:         err.Error(),
				SupportedSounds: sounds.SupportedSounds,
			})
		}
		// If there was an issue with assistant relay
		if resp.StatusCode() != http.StatusOK {
			return c.JSON(http.StatusInternalServerError, SoundFailedPlayResponse{
				Status:          failedString,
				Message:         resp.String(),
				SupportedSounds: sounds.SupportedSounds,
			})
		}
		// If sound not found in support sounds
	} else {
		return c.JSON(http.StatusBadRequest, SoundFailedPlayResponse{
			Status:          failedString,
			Message:         fmt.Sprintf("sound: %v not supported", sound),
			SupportedSounds: sounds.SupportedSounds,
		})
	}
	return c.JSON(http.StatusOK, &SoundSuccessfulPlayResponse{
		Status: successString,
	})
}
