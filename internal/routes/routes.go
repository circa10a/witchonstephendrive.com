package routes

import (
	"io/fs"
	"net/http"

	_ "github.com/circa10a/witchonstephendrive.com/api" // import generated docs.go
	"github.com/circa10a/witchonstephendrive.com/internal/colors"
	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	swagger "github.com/swaggo/echo-swagger"
)

const (
	successString = "success"
	failedString  = "failed"
)

// Routes instantiates all of the listening context paths
func Routes(e *echo.Echo, witchConfig config.WitchConfig, frontendAssets fs.FS, apiDocAssets fs.FS) {
	// Static assets
	frontendHTTPFS, err := utils.ConvertEmbedFsDirToHTTPFS(frontendAssets, "web")
	if err != nil {
		log.Error(err)
	}
	// Frontend html/css/js
	frontendFileServer := http.FileServer(frontendHTTPFS)
	// Serve frontend static assets
	e.GET("/*", echo.WrapHandler(frontendFileServer))
	// Swagger.{json,yaml}
	apiDocsFileServer := http.FileServer(http.FS(apiDocAssets))

	// API docs/Swagger JSON
	e.GET("/api/*", echo.WrapHandler(apiDocsFileServer))

	// Lights
	// Route to view supported colors
	e.GET("/colors", colorsReadHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})
	// Route to change color of lights
	e.POST("/color/:color", colorChangeHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		// In the event user passes unsupported color, give them a list
		return func(c echo.Context) error {
			c.Set("hueLights", witchConfig.HueLights)
			c.Set("bridge", witchConfig.Bridge)
			c.Set("supportedColors", colors.SupportedColors)
			return next(c)
		}
	})

	// Route to change lights state(on/off)
	e.POST("/lights/:state", lightsStateHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("hueLights", witchConfig.HueLightsStructs)
			return next(c)
		}
	})

	// Sounds
	// Route to view supported sounds
	e.GET("/sounds", soundsReadHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})
	// Route to play sounds
	e.POST("/sound/:sound", soundPlayHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("client", witchConfig.Client)
			c.Set("assistantDevice", witchConfig.AssistantDevice)
			return next(c)
		}
	})

	// Swagger docs
	url := swagger.URL("/api/swagger.json")
	e.GET("/swagger/*", swagger.EchoWrapHandler(url))
}
