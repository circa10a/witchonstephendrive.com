package routes

import (
	"io/fs"
	"net/http"

	_ "github.com/circa10a/witchonstephendrive.com/api" // import generated docs.go
	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	"github.com/circa10a/witchonstephendrive.com/routes/handlers"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	swagger "github.com/swaggo/echo-swagger"
)

// Routes instantiates all of the listening context paths
func Routes(e *echo.Echo, witchConfig *config.WitchConfig, frontendAssets fs.FS, apiDocAssets fs.FS) {
	apiVersionGroup := e.Group(witchConfig.APIBaseURL)
	// UI
	// Static assets
	if witchConfig.UIEnabled {
		frontendHTTPFS, err := utils.ConvertEmbedFsDirToHTTPFS(frontendAssets, "web")
		if err != nil {
			log.Error(err)
		}
		// Frontend html/css/js
		frontendFileServer := http.FileServer(frontendHTTPFS)
		// Serve frontend static assets
		e.GET("/*", echo.WrapHandler(frontendFileServer))
	}

	// API
	// Swagger.{json,yaml}
	apiDocsFileServer := http.FileServer(http.FS(apiDocAssets))
	// API docs/Swagger JSON
	e.GET("/api/*", echo.WrapHandler(apiDocsFileServer))
	// Swagger docs
	swaggerURL := swagger.URL("/api/swagger.json")
	e.GET("/swagger/*", swagger.EchoWrapHandler(swaggerURL))

	// Lights/Colors
	// Route to view supported colors
	apiVersionGroup.GET("/colors", handlers.ColorsReadHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})
	// Route to change color of lights
	apiVersionGroup.POST("/color/:color", handlers.ColorChangeHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		// In the event user passes unsupported color, give them a list
		return func(c echo.Context) error {
			c.Set("hueLights", witchConfig.HueLightsStructs)
			c.Set("hueBridge", witchConfig.HueBridge)
			return next(c)
		}
	})

	// Route to change lights state(on/off)
	apiVersionGroup.POST("/lights/:state", handlers.LightsStateHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("hueLights", witchConfig.HueLightsStructs)
			return next(c)
		}
	})

	// Sounds
	// Route to view supported sounds
	apiVersionGroup.GET("/sounds", handlers.SoundsReadHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})
	// Route to play sounds
	apiVersionGroup.POST("/sound/:sound", handlers.SoundPlayHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Route is only functional if an entity ID is configured
			c.Set("homeAssistantEntityID", witchConfig.HomeAssistantEntityID)
			c.Set("quietTimeEnabled", witchConfig.SoundQuietTimeEnabled)
			c.Set("quietTimeStart", witchConfig.SoundQuietTimeStart)
			c.Set("quietTimeEnd", witchConfig.SoundQuietTimeEnd)
			c.Set("soundQueue", witchConfig.SoundQueue)
			return next(c)
		}
	})
}
