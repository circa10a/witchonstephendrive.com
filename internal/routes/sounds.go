package routes

import (
	"fmt"
	"net/http"

	"github.com/circa10a/witchonstephendrive.com/internal/sounds"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/oleiade/lane"
)

// SoundsListResponse responds supported sounds to play
type SoundsListResponse struct {
	SupportedSounds []string `json:"supportedSounds"`
}

// SoundSuccessfulPlayResponse responds with a boolean to indicate successful or not
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
// @Success 202 {object} SoundSuccessfulPlayResponse
// @Failure 400 {object} SoundFailedPlayResponse
// @Router /sound/{sound} [post]
// @Param sound path string true "Sound to play"
func soundPlayHandler(c echo.Context) error {
	sound := c.Param("sound")
	queue := c.Get("soundQueue").(*lane.Queue)
	quietTimeStart := c.Get("quietTimeStart").(string)
	quietTimeEnd := c.Get("quietTimeEnd").(string)
	// Ensure sounds don't play during quiet time(late hours)
	if sounds.IsDuringQuietTime(quietTimeStart, quietTimeEnd) {
		return c.JSON(http.StatusBadRequest, SoundFailedPlayResponse{
			Success:         false,
			Message:         fmt.Sprintf("quiet time enabled. quiet time is between %s and %s", quietTimeStart, quietTimeEnd),
			SupportedSounds: sounds.SupportedSounds,
		})
	}
	if utils.StrInSlice(sound, sounds.SupportedSounds) {
		queue.Enqueue(sound)
	} else {
		// If sound not found in supported sounds
		return c.JSON(http.StatusBadRequest, SoundFailedPlayResponse{
			Success:         false,
			Message:         fmt.Sprintf("sound: %v not supported", sound),
			SupportedSounds: sounds.SupportedSounds,
		})
	}
	return c.JSON(http.StatusAccepted, &SoundSuccessfulPlayResponse{
		Success: true,
	})
}
