package routes

import (
	"fmt"
	"net/http"

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
