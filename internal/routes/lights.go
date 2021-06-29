package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/internal/colors"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// ColorsListResponse responds supported colors to set
type ColorsListResponse struct {
	SupportedColors []string `json:"supportedColors"`
}

// ColorSuccessfulChangeResponse responds with a success status string when operation has completed successfully
type ColorSuccessfulChangeResponse struct {
	Status string `json:"status"`
}

// ColorFailedChangeResponse responds with status string, reason for failure in message, and list of supported colors
type ColorFailedChangeResponse struct {
	Status          string   `json:"status"`
	Message         string   `json:"message"`
	SupportedColors []string `json:"supportedColors"`
}

// LightStateSuccessfulChangeResponse responds with status string
type LightStateSuccessfulChangeResponse struct {
	Status string `json:"status"`
}

// LightStateFailedChangeResponse responds with status string, reason for failure in message
type LightStateFailedChangeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// colors godoc
// @Summary Get available colors to change to
// @Description Get list of supported colors
// @Produce json
// @Success 200 {object} ColorsListResponse
// @Router /colors [get]
func colorsReadHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, ColorsListResponse{
		SupportedColors: colors.SupportedColors,
	})
}

// :color godoc
// @Summary Change hue lights color
// @Description Change hue lights to supported color defined in color map
// @Produce json
// @Success 200 {object} ColorSuccessfulChangeResponse
// @Failure 400 {object} ColorFailedChangeResponse
// @Failure 500 {object} ColorFailedChangeResponse
// @Router /color/{color} [post]
// @Param color path string true "Color to change lights to"
func colorChangeHandler(c echo.Context) error {
	colorsMap := colors.Colors
	color := c.Param("color")
	hueLights := c.Get("hueLights").([]int)
	bridge := c.Get("bridge").(*huego.Bridge)
	for _, light := range hueLights {
		// Only change color if in the map
		if _, ok := colorsMap[color]; ok {
			_, err := bridge.SetLightState(light, colorsMap[color])
			if err != nil {
				log.Error(err)
				return c.JSON(http.StatusInternalServerError, &ColorFailedChangeResponse{
					Status:          failedString,
					Message:         err.Error(),
					SupportedColors: colors.SupportedColors,
				})
			}
			// Fail if color not supported
		} else {
			return c.JSON(http.StatusBadRequest, &ColorFailedChangeResponse{
				Status:          failedString,
				Message:         fmt.Sprintf("color: %v not supported", color),
				SupportedColors: colors.SupportedColors,
			})
		}
	}
	return c.JSON(http.StatusOK, &ColorSuccessfulChangeResponse{
		Status: successString,
	})
}

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
			Status:  failedString,
			Message: fmt.Sprintf("received state: %v. on/off are the only valid values", state),
		})
	}
	// Loop through lights and change state accordingly
	for _, light := range hueLights {
		if state == "on" {
			err := light.On()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, LightStateFailedChangeResponse{
					Status:  failedString,
					Message: err.Error(),
				})
			}
		}
		if state == "off" {
			err := light.Off()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, LightStateFailedChangeResponse{
					Status:  failedString,
					Message: err.Error(),
				})
			}
		}
	}
	return c.JSON(http.StatusOK, LightStateSuccessfulChangeResponse{
		Status: successString,
	})
}
