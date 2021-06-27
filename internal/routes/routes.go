package routes

import (
	"io/fs"
	"net/http"

	"github.com/amimof/huego"
	_ "github.com/circa10a/witchonstephendrive.com/api" // import generated docs.go
	"github.com/circa10a/witchonstephendrive.com/internal/colors"
	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	swagger "github.com/swaggo/echo-swagger"
)

// ColorResponse responds with success or failed status string
type ColorResponse struct {
	Status string
}

// SoundResponse responds with success or failed status string
type SoundResponse struct {
	Status string
}

// :color godoc
// @Summary Change hue lights color
// @Description Change hue lights to supported color defined in color map
// @Produce json
// @Success 200 {object} ColorResponse
// @Failure 400 {object} ColorResponse
// @Router /color/{color} [post]
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

// :sound godoc
// @Summary Play sound via assistant relay
// @Description Play halloween sound supported sound defined in sound map
// @Produce json
// @Success 200 {object} SoundResponse
// @Failure 400 {object} SoundResponse
// @Router /sound/{sound} [post]
// @Param sound path string true "Sound to play"
func soundHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, SoundResponse{
		Status: "success",
	})
}

// Routes instantiates all of the listening context paths
func Routes(e *echo.Echo, witchConfig config.WitchConfig, frontendAssets fs.FS, apiDocAssets fs.FS) {
	frontendHTTPFS, err := utils.ConvertEmbedFsDirToHTTPFS(frontendAssets, "web")
	if err != nil {
		log.Error(err)
	}

	frontendFileServer := http.FileServer(frontendHTTPFS)
	// Serve frontend static assets
	e.GET("/*", echo.WrapHandler(frontendFileServer))

	apiDocsFileServer := http.FileServer(http.FS(apiDocAssets))
	// API docs/Swagger JSON
	e.GET("/api/*", echo.WrapHandler(apiDocsFileServer))

	// Share huelight id's and bridge connection with route group
	// Route to change lights
	e.POST("/color/:color", colorHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("hueLights", witchConfig.HueLights)
			c.Set("bridge", witchConfig.Bridge)
			return next(c)
		}
	})

	e.POST("/sound/:sound", soundHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})

	// Swagger docs
	url := swagger.URL("/api/swagger.json")
	e.GET("/swagger/*", swagger.EchoWrapHandler(url))
}
