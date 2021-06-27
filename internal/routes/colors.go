package routes

import (
	"fmt"
	"net/http"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/internal/colors"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// ColorResponse responds with success or failed status string and list of supported colors
type ColorResponse struct {
	Status          string   `json:"status"`
	Message         string   `json:"message,omitempty"`
	SupportedColors []string `json:"supportedColors"`
}

// colors godoc
// @Summary Get available colors to change to
// @Description Get list of supported colors
// @Produce json
// @Success 200 {object} ColorResponse
// @Failure 400 {object} ColorResponse
// @Router /colors [get]
func colorsReadHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, ColorResponse{
		Status:          successString,
		SupportedColors: colors.SupportedColors,
	})
}

// :color godoc
// @Summary Change hue lights color
// @Description Change hue lights to supported color defined in color map
// @Produce json
// @Success 200 {object} ColorResponse
// @Failure 400 {object} ColorResponse
// @Router /color/{color} [post]
// @Param color path string true "Color to change lights to"
func colorChangeHandler(c echo.Context) error {
	colorsMap := colors.Colors
	color := c.Param("color")
	hueLights := c.Get("hueLights").([]int)
	bridge := c.Get("bridge").(*huego.Bridge)
	responseCode := http.StatusOK
	response := &ColorResponse{
		Status:          successString,
		SupportedColors: colors.SupportedColors,
	}
	for _, light := range hueLights {
		// Only change color if in the map
		if _, ok := colorsMap[color]; ok {
			_, err := bridge.SetLightState(light, colorsMap[color])
			if err != nil {
				log.Error(err)
				responseCode = http.StatusBadRequest
				response.Status = failedString
				response.Message = err.Error()
			}
			// Fail if color not supported
		} else {
			responseCode = http.StatusBadRequest
			response.Status = failedString
			response.Message = fmt.Sprintf("color: %v not supported", color)
		}
	}
	return c.JSON(responseCode, response)
}
