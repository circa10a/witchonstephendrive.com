package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/amimof/huego"
	"github.com/labstack/echo/v4"
)

// :state godoc
// @Summary Change state of configured lights
// @Description Only supports states of on/off
// @Produce json
// @Success 200 {object} LightStateSuccessfulChangeResponse
// @Failure 400 {object} LightStateFailedChangeResponse
// @Failure 500 {object} LightStateFailedChangeResponse
// @Router /lights/{state} [post]
// @Param state path string true "State to set lights to (on/off)"
func lightsStateHandler(c echo.Context) error {
	state := strings.ToLower(c.Param("state"))
	hueLights := c.Get("hueLights").([]huego.Light)

	// Check for on/off states
	if state != "on" && state != "off" {
		return c.JSON(http.StatusBadRequest, LightStateFailedChangeResponse{
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
			return c.JSON(http.StatusInternalServerError, LightStateFailedChangeResponse{
				Success: false,
				Message: err.Error(),
			})
		}
	}
	return c.JSON(http.StatusOK, LightStateSuccessfulChangeResponse{
		Success: true,
	})
}
