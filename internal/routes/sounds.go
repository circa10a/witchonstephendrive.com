package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/circa10a/witchonstephendrive.com/internal/sounds"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	"github.com/go-redis/redis/v8"
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
// @Failure 500 {object} SoundFailedPlayResponse
// @Router /sound/{sound} [post]
// @Param sound path string true "Sound to play"
func soundPlayHandler(c echo.Context) error {
	sound := c.Param("sound")
	redisChannel := c.Get("redisChannel").(string)
	redisClient := c.Get("redisClient").(*redis.Client)
	redisContext := c.Get("redisContext").(context.Context)
	if utils.StrInSlice(sound, sounds.SupportedSounds) {
		// Drop message to redis queue to be played
		err := redisClient.Publish(redisContext, redisChannel, sound).Err()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, SoundFailedPlayResponse{
				Success:         false,
				Message:         err.Error(),
				SupportedSounds: sounds.SupportedSounds,
			})
		}
		// If sound not found in supported sounds
	} else {
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
