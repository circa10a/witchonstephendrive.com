package routes

import (
	"fmt"
	"net/http"

	"github.com/circa10a/witchonstephendrive.com/internal/sounds"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	"github.com/labstack/echo/v4"
)

// SoundsListResponse responds supported sounds to play
type SoundsListResponse struct {
	SupportedSounds []string `json:"supportedSounds"`
}

// SoundSuccesfulPlayResponse responds with a boolean to indicate successful or not
type SoundSuccessfulPlayResponse struct {
	Success bool `json:"success"`
}

// SoundFailedPlayResponse responds with a boolean to indicate successful or not,
// reason for failure, and list of supported sounds
type SoundFailedPlayResponse struct {
	Success         bool     `json:"success"`
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
// @Router /sound/{sound} [post]
// @Param sound path string true "Sound to play"
func soundPlayHandler(c echo.Context) error {
	sound := c.Param("sound")
	channel := c.Get("soundChannel").(chan string)
	if utils.StrInSlice(sound, sounds.SupportedSounds) {
		channel <- sound
	} else {
		// If sound not found in supported sounds
		return c.JSON(http.StatusBadRequest, SoundFailedPlayResponse{
			Success:         false,
			Message:         fmt.Sprintf("sound: %v not supported", sound),
			SupportedSounds: sounds.SupportedSounds,
		})
	}
	return c.JSON(http.StatusOK, &SoundSuccessfulPlayResponse{
		Success: true,
	})
}
