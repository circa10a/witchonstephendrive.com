package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/controllers/colors"
	"github.com/labstack/echo/v4"
)

// ColorsListResponse responds supported colors to set
type ColorsListResponse struct {
	SupportedColors []string `json:"supportedColors"`
}

// ColorChangeResponse responds with a boolean to indicate successful or not and message
type ColorChangeResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// colors godoc
// @Summary Get available colors to change to
// @Description Get list of supported colors
// @Produce json
// @Success 200 {object} ColorsListResponse
// @Router /colors [get]
func ColorsReadHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, ColorsListResponse{
		SupportedColors: colors.SupportedColors,
	})
}

// :color godoc
// @Summary Change hue lights color
// @Description Change hue lights to supported color defined in color map
// @Produce json
// @Success 200 {object} ColorChangeResponse
// @Failure 400 {object} ColorChangeResponse
// @Failure 403 {object} ColorChangeResponse
// @Failure 500 {object} ColorChangeResponse
// @Router /color/{color} [post]
// @Param color path string true "Color to change lights to"
func ColorChangeHandler(c echo.Context) error {
	color := c.Param("color")
	hueLights := c.Get("hueLights").([]huego.Light)
	hueBridge := c.Get("hueBridge").(*huego.Bridge)
	errs := colors.SetLightsColor(hueLights, hueBridge, color)
	// Send back first error if exists
	if len(errs) > 0 {
		if errors.Is(errs[0], colors.ErrColorNotSupported) {
			return c.JSON(http.StatusBadRequest, &ColorChangeResponse{
				Success: false,
				Message: fmt.Sprintf("color: %v not supported", color),
			})
		} else {
			return c.JSON(http.StatusInternalServerError, &ColorChangeResponse{
				Success: false,
				Message: errs[0].Error(),
			})
		}
	}
	return c.JSON(http.StatusOK, &ColorChangeResponse{
		Success: true,
		Message: fmt.Sprintf("color: %s set successfully", color),
	})
}
