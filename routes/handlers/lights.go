package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/amimof/huego"
	"github.com/labstack/echo/v4"
)

// LightStateChangeResponse responds with a boolean to indicate successful or not and message
type LightStateChangeResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// :state godoc
// @Summary Change state of configured lights
// @Description Only supports states of on/off
// @Produce json
// @Success 200 {object} LightStateChangeResponse
// @Failure 400 {object} LightStateChangeResponse
// @Failure 500 {object} LightStateChangeResponse
// @Router /lights/{state} [post]
// @Param state path string true "State to set lights to (on/off)"
func LightsStateHandler(c echo.Context) error {
	state := strings.ToLower(c.Param("state"))
	hueLights := c.Get("hueLights").([]huego.Light)

	// Check for on/off states
	if state != "on" && state != "off" {
		return c.JSON(http.StatusBadRequest, LightStateChangeResponse{
			Success: false,
			Message: fmt.Sprintf("received state: %v. on/off are the only valid values", state),
		})
	}

	// Loop through lights and change state accordingly
	for _, light := range hueLights {
		var err error
		if state == "on" {
			err = light.On()
		} else if state == "off" {
			err = light.Off()
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, LightStateChangeResponse{
				Success: false,
				Message: err.Error(),
			})
		}
	}
	return c.JSON(http.StatusOK, LightStateChangeResponse{
		Success: true,
		Message: fmt.Sprintf("light state: %s set successfully", state),
	})
}