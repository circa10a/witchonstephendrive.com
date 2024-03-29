package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/controllers/lights"
	"github.com/labstack/echo/v4"
)

// LightStateChangeResponse responds with a boolean to indicate successful or not and message
type LightStateChangeResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// PostLightsHandler holds all the data needed for the POST lights handler
type PostLightsHandler struct {
	echo.Context
	HueLights []huego.Light
}

// :state godoc
// @Summary Change state of configured lights
// @Description Only supports states of on/off
// @Produce json
// @Success 200 {object} LightStateChangeResponse
// @Failure 400 {object} LightStateChangeResponse
// @Failure 403 {object} LightStateChangeResponse
// @Failure 500 {object} LightStateChangeResponse
// @Router /lights/{state} [post]
// @Param state path string true "State to set lights to (on/off)"
func (h *PostLightsHandler) Handler(c echo.Context) error {
	state := strings.ToLower(c.Param("state"))

	// Check for on/off states
	if state != "on" && state != "off" {
		return c.JSON(http.StatusBadRequest, LightStateChangeResponse{
			Success: false,
			Message: fmt.Sprintf("received state: %v. on/off are the only valid values", state),
		})
	}

	var err error
	if state == "on" {
		err = lights.SetLightsOn(h.HueLights)
	} else if state == "off" {
		err = lights.SetLightsOff(h.HueLights)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, LightStateChangeResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, LightStateChangeResponse{
		Success: true,
		Message: fmt.Sprintf("light state: %s set successfully", state),
	})
}
