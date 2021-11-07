package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/circa10a/witchonstephendrive.com/controllers/sounds"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/oleiade/lane"
)

// SoundsListResponse responds supported sounds to play
type SoundsListResponse struct {
	SupportedSounds []string `json:"supportedSounds"`
}

// SoundPlayResponse responds with a boolean to indicate successful or not and message
type SoundPlayResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// sounds godoc
// @Summary Get available sounds to play
// @Description Get list of supported sounds
// @Produce json
// @Success 200 {object} SoundsListResponse
// @Router /sounds [get]
func SoundsReadHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, SoundsListResponse{
		SupportedSounds: sounds.SupportedSounds,
	})
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
func SoundPlayHandler(c echo.Context) error {
	sound := c.Param("sound")
	homeAssistantEntityID := c.Get("homeAssistantEntityID").(string)
	queue := c.Get("soundQueue").(*lane.Deque)
	quietTimeEnabled := c.Get("quietTimeEnabled").(bool)
	quietTimeStart := c.Get("quietTimeStart").(int)
	quietTimeEnd := c.Get("quietTimeEnd").(int)
	// Only enable sounds if entity ID is configured
	// The ensures sounds never reach the queue
	if homeAssistantEntityID == "" {
		return c.JSON(http.StatusBadRequest, SoundPlayResponse{
			Success: false,
			Message: "sounds disabled. no home assistant entity ID configured",
		})
	}
	// Ensure sounds don't play during quiet time(late hours)
	if sounds.IsDuringQuietTime(time.Now().Hour(), quietTimeStart, quietTimeEnd) && quietTimeEnabled {
		return c.JSON(http.StatusBadRequest, SoundPlayResponse{
			Success: false,
			Message: fmt.Sprintf("sounds disabled. quiet time is between %d:00 and %d:00", quietTimeStart, quietTimeEnd),
		})
	}
	// If sound is supported
	if utils.StrInSlice(sound, sounds.SupportedSounds) {
		// Ensure we don't get a huge backlog of sound requests by limiting with a capped queue
		if !queue.Append(sound) {
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
