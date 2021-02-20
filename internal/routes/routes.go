package routes

import (
	"net/http"

	"github.com/amimof/huego"
	_ "github.com/circa10a/witchonstephendrive.com/api" // import generated docs.go
	"github.com/circa10a/witchonstephendrive.com/internal/colors"
	swagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// ColorResponse responds with success or failed status string
type ColorResponse struct {
	Status string
}

// :color godoc
// @Summary Change hue lights color
// @Description Change hue lights to supported color defined in color map
// @Produce json
// @Success 200 {object} ColorResponse
// @Failure 400 {object} ColorResponse
// @Router /{color} [post]
// @Param color path string true "Color to change lights to"
func colorHandler(c echo.Context) error {
	colors := colors.Colors
	color := c.Param("color")
	hueLights := c.Get("hueLights").([]int)
	bridge := c.Get("bridge").(*huego.Bridge)
	for _, light := range hueLights {
		// Only change color if in the map
		if _, ok := colors[color]; ok {
			_, err := bridge.SetLightState(light, colors[color])
			if err != nil {
				log.Error(err)
			}
		} else {
			return c.JSON(http.StatusBadRequest, ColorResponse{
				Status: "failed",
			})
		}
	}
	return c.JSON(http.StatusOK, ColorResponse{
		Status: "success",
	})
}

// Routes instantiates all of the listening context paths
func Routes(e *echo.Echo, hueLights []int, bridge *huego.Bridge) {
	// Serve frontend static assets
	e.Static("/", "web")
	// Swagger docs
	e.Static("/api", "./api")
	// Share huelight id's and bridge connection with route group
	// Route to change lights
	e.POST("/api/:color", colorHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("hueLights", hueLights)
			c.Set("bridge", bridge)
			return next(c)
		}
	})
	e.GET("/swagger/*", swagger.WrapHandler)
}
