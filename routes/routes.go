package routes

import (
	"io/fs"
	"net/http"

	_ "github.com/circa10a/witchonstephendrive.com/api" // import generated docs.go
	"github.com/circa10a/witchonstephendrive.com/controllers/colors"
	"github.com/circa10a/witchonstephendrive.com/controllers/sounds"
	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/routes/handlers"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	swagger "github.com/swaggo/echo-swagger"
)

const (
	colorsPath      = "/colors"
	colorNamePath   = "/color/:color"
	lightStatePath  = "/lights/:state"
	soundsPath      = "/sounds"
	soundNamePath   = "/sound/:sound"
	swaggerDocsPath = "/swagger/*"
	apiDocsPath     = "/api/*"
	frontendPath    = "/*"
)

// Routes instantiates all of the listening context paths
func Routes(e *echo.Echo, witchConfig *config.WitchConfig, frontendAssets fs.FS, apiDocAssets fs.FS) {
	apiVersionGroup := e.Group(witchConfig.APIBaseURL)
	// UI
	// Static assets
	if witchConfig.UIEnabled {
		frontendHTTPFS, err := ConvertEmbedFsDirToHTTPFS(frontendAssets, "web")
		if err != nil {
			log.Error(err)
		}
		// Frontend html/css/js
		frontendFileServer := http.FileServer(frontendHTTPFS)
		// Serve frontend static assets
		e.GET(frontendPath, echo.WrapHandler(frontendFileServer))
	}

	// API
	// Swagger.{json,yaml}
	apiDocsFileServer := http.FileServer(http.FS(apiDocAssets))
	// API docs/Swagger JSON
	e.GET(apiDocsPath, echo.WrapHandler(apiDocsFileServer))

	// Swagger docs
	swaggerURL := swagger.URL("/api/swagger.json")
	e.GET(swaggerDocsPath, swagger.EchoWrapHandler(swaggerURL))

	// Lights/Colors
	// Route to view supported colors
	getColorsHandler := &handlers.GetColorsHandler{
		SupportedColors: colors.SupportedColors,
	}
	apiVersionGroup.GET(colorsPath, getColorsHandler.Handler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})

	postColorsHandler := &handlers.PostColorsHandler{
		HueBridge: witchConfig.HueBridge,
		HueLights: witchConfig.HueLightsStructs,
	}
	// Route to change color of lights
	apiVersionGroup.POST(colorNamePath, postColorsHandler.Handler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})

	postLightsHandler := &handlers.PostLightsHandler{
		HueLights: witchConfig.HueLightsStructs,
	}
	// Route to change lights state(on/off)
	apiVersionGroup.POST(lightStatePath, postLightsHandler.Handler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})

	// Sounds
	// Route to view supported sounds
	getSoundsHandler := &handlers.GetSoundsHandler{
		SupportedSounds: sounds.SupportedSounds,
	}
	apiVersionGroup.GET(soundsPath, getSoundsHandler.Handler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})
	// Route to play sounds
	postSoundsHandler := &handlers.PostSoundsHandler{
		Queue:                 witchConfig.SoundQueue,
		QuietTimeEnabled:      witchConfig.SoundQuietTimeEnabled,
		QuietTimeStart:        witchConfig.SoundQuietTimeStart,
		QuietTimeEnd:          witchConfig.SoundQuietTimeEnd,
		HomeAssistantEntityID: witchConfig.HomeAssistantEntityID,
	}
	apiVersionGroup.POST(soundNamePath, postSoundsHandler.Handler, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Route is only functional if an entity ID is configured
			return next(c)
		}
	})
}

// ConvertEmbedFsDirToHTTPSFS returns sub directory of fs
func ConvertEmbedFsDirToHTTPFS(e fs.FS, d string) (http.FileSystem, error) {
	fsys, err := fs.Sub(e, d)
	if err != nil {
		return nil, err
	}
	return http.FS(fsys), nil
}
