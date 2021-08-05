package routes

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/internal/colors"
	"github.com/labstack/echo/v4"
)

// ColorsListResponse responds supported colors to set
type ColorsListResponse struct {
	SupportedColors []string `json:"supportedColors"`
}

// ColorSuccessfulChangeResponse responds with a boolean to indicate successful or not
type ColorSuccessfulChangeResponse struct {
	Success bool `json:"success"`
}

// ColorFailedChangeResponse responds with status string, reason for failure, and list of supported colors
type ColorFailedChangeResponse struct {
	Success         bool     `json:"success"`
	Message         string   `json:"message"`
	SupportedColors []string `json:"supportedColors"`
}

// LightStateSuccessfulChangeResponse responds with a boolean to indicate successful or not
type LightStateSuccessfulChangeResponse struct {
	Success bool `json:"success"`
}

// LightStateFailedChangeResponse responds with a boolean to indicate successful or not and reason for failure
type LightStateFailedChangeResponse struct {
	Success bool   `json:"success"`
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
	color := c.Param("color")
	hueLights := c.Get("hueLights").([]int)
	hueBridge := c.Get("hueBridge").(*huego.Bridge)
	err := colors.SetLightsColor(hueLights, hueBridge, color)
	if err != nil {
		if errors.Is(err, colors.ErrColorNotSupported) {
			return c.JSON(http.StatusBadRequest, &ColorFailedChangeResponse{
				Success:         false,
				Message:         fmt.Sprintf("color: %v not supported", color),
				SupportedColors: colors.SupportedColors,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, &ColorFailedChangeResponse{
				Success:         false,
				Message:         err.Error(),
				SupportedColors: colors.SupportedColors,
			})
		}
	}
	return c.JSON(http.StatusOK, &ColorSuccessfulChangeResponse{
		Success: true,
	})
}
