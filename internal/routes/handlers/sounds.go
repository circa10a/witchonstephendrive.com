package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/circa10a/witchonstephendrive.com/controllers/sounds"
	"github.com/labstack/echo/v4"
)

// SoundsListResponse responds supported sounds to play
type SoundsListResponse struct {
	SupportedSounds []string `json:"supportedSounds"`
}

// SoundPlayResponse responds with a boolean to indicate successful or not and message
type SoundPlayResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// GetSoundsHandler holds all the data needed for the GET sounds handler
type GetSoundsHandler struct {
	echo.Context
	SupportedSounds []string
}

// sounds godoc
// @Summary Get available sounds to play
// @Description Get list of supported sounds
// @Produce json
// @Success 200 {object} SoundsListResponse
// @Router /sounds [get]
func (h *GetSoundsHandler) Handler(c echo.Context) error {
	return c.JSON(http.StatusOK, SoundsListResponse{
		SupportedSounds: h.SupportedSounds,
	})
}

// PostSoundsHandler holds all the data needed for the POST sounds handler
type PostSoundsHandler struct {
	echo.Context
	Queue                 chan string
	HomeAssistantEntityID string
	QuietTimeStart        int
	QuietTimeEnd          int
	QuietTimeEnabled      bool
}

// :sound godoc
// @Summary Play sound via home assistant
// @Description Play halloween sound supported in sound list
// @Produce json
// @Success 202 {object} SoundPlayResponse
// @Failure 400 {object} SoundPlayResponse
// @Failure 403 {object} SoundPlayResponse
// @Failure 429 {object} SoundPlayResponse
// @Router /sound/{sound} [post]
// @Param sound path string true "Sound to play"
func (h *PostSoundsHandler) Handler(c echo.Context) error {
	sound := c.Param("sound")
	// Only enable sounds if entity ID is configured
	// The ensures sounds never reach the queue
	if h.HomeAssistantEntityID == "" {
		return c.JSON(http.StatusBadRequest, SoundPlayResponse{
			Success: false,
			Message: "sounds disabled. no home assistant entity ID configured",
		})
	}
	// Ensure sounds don't play during quiet time(late hours)
	if sounds.IsDuringQuietTime(time.Now().Hour(), h.QuietTimeStart, h.QuietTimeEnd) && h.QuietTimeEnabled {
		return c.JSON(http.StatusBadRequest, SoundPlayResponse{
			Success: false,
			Message: fmt.Sprintf("sounds disabled. quiet time is between %d:00 and %d:00", h.QuietTimeStart, h.QuietTimeEnd),
		})
	}
	// If sound is supported
	if StrInSlice(sound, sounds.SupportedSounds) {
		// Ensure we don't get a huge backlog of sound requests by limiting with a capped queue
		// If queue is at capacity, enter default case
		select {
		case h.Queue <- sound:
			// Playing sound
		default:
			return c.JSON(http.StatusTooManyRequests, SoundPlayResponse{
				Success: false,
				Message: fmt.Sprintf("will not play sound: %s. sound queue is at capacity", sound),
			})
		}
	} else {
		// If sound not found in supported sounds
		return c.JSON(http.StatusBadRequest, SoundPlayResponse{
			Success: false,
			Message: fmt.Sprintf("sound: %v not supported", sound),
		})
	}
	// This means queue.append(sound) was successful
	return c.JSON(http.StatusAccepted, &SoundPlayResponse{
		Success: true,
		Message: fmt.Sprintf("sound: %s queued successfully", sound),
	})
}

// StrInSlice returns true if string is in slice
func StrInSlice(str string, list []string) bool {
	for _, item := range list {
		if str == item {
			return true
		}
	}
	return false
}
