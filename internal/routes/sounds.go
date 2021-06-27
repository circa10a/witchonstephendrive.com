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

// SoundResponse responds with success or failed status string
type SoundResponse struct {
	Status          string   `json:"status"`
	Message         string   `json:"message,omitempty"`
	SupportedSounds []string `json:"supportedSounds"`
}

// sounds godoc
// @Summary Get available sounds to play
// @Description Get list of supported sounds
// @Produce json
// @Success 200 {object} SoundResponse
// @Failure 400 {object} SoundResponse
// @Router /sounds [get]
func soundsReadHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, SoundResponse{
		Status:          successString,
		SupportedSounds: sounds.SupportedSounds,
	})
}

// :sound godoc
// @Summary Play sound via assistant relay
// @Description Play halloween sound supported in sound list
// @Produce json
// @Success 200 {object} SoundResponse
// @Failure 400 {object} SoundResponse
// @Router /sound/{sound} [post]
// @Param sound path string true "Sound to play"
func soundPlayHandler(c echo.Context) error {
	sound := c.Param("sound")
	client := c.Get("client").(*resty.Client)
	assistantDevice := c.Get("assistantDevice")
	responseCode := http.StatusOK
	response := &SoundResponse{
		Status:          successString,
		SupportedSounds: sounds.SupportedSounds,
	}
	if utils.StrInSlice(sound, sounds.SupportedSounds) {
		// Call assistant-relay if sound is supported
		resp, err := client.R().SetBody(&sounds.PlaySoundPayload{
			Device: assistantDevice.(string),
			Source: fmt.Sprintf("%v.mp3", sound),
			Type:   "custom",
		}).Post("/cast")
		// Handle unknown error
		if err != nil || resp.StatusCode() != http.StatusOK {
			log.Error(err)
			responseCode = http.StatusBadRequest
			response.Status = failedString
			response.Message = err.Error()
		}
	} else {
		// If sound is not supported
		responseCode = http.StatusBadRequest
		response.Status = failedString
		response.Message = fmt.Sprintf("sound: %v not supported", sound)
	}
	return c.JSON(responseCode, response)
}
