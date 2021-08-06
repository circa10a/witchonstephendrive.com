package main

import (
	"embed"
	"fmt"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/controllers/sounds"
	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/routes"
	witchPrometheusMiddleware "github.com/circa10a/witchonstephendrive.com/routes/middleware/prometheus"
	"github.com/go-resty/resty/v2"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oleiade/lane"
	log "github.com/sirupsen/logrus"
)

//go:embed web
var frontendAssets embed.FS

//go:embed api
var apiDocAssets embed.FS

// @title witchonstephendrive.com
// @version 0.1.0
// @description Control my halloween decorations
// @contact.name Caleb Lemoine
// @contact.email caleblemoine@gmail.com
// @contact.url https://caleblemoine.dev
// @license.name MIT
// @license.url https://raw.githubusercontent.com/circa10a/witchonstephendrive.com/main/LICENSE
// @tag.Name
// @tag.description
// @tag.docs.url https://github.com/circa10a/witchonstephendrive.com
// @tag.docs.description Link to GitHub Repository
// @host witchonstephendrive.com
// @BasePath /api/v1
// @Schemes https
func main() {
	// Setup global config store
	witchConfig := config.WitchConfig{}
	err := envconfig.Process("", &witchConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Hue Lights Config
	// Find hue bridge ip
	hueBridge, err := huego.Discover()
	if err != nil {
		log.Fatal(err)
	}
	witchConfig.HueBridge = hueBridge
	// Authenticate against bridge api
	witchConfig.HueBridge.Login(witchConfig.HueUser)
	// Store all light data to be used later
	for _, lightID := range witchConfig.HueLights {
		light, err := witchConfig.HueBridge.GetLight(lightID)
		if err != nil {
			log.Fatal(err)
		}
		witchConfig.HueLightsStructs = append(witchConfig.HueLightsStructs, *light)
	}

	// Create client to be used with assistant relay
	assitantRelayEndpoint := fmt.Sprintf("%s:%d", witchConfig.AssistantRelayHost, witchConfig.AssistantRelayPort)
	witchConfig.RelayClient = resty.New().SetHostURL(assitantRelayEndpoint).SetHeader("Content-Type", "application/json")

	// Sound processing
	// Create new queue to process "sound" jobs one at a time with a max limit to eliminate spam
	witchConfig.SoundQueue = lane.NewCappedDeque(witchConfig.SoundQueueCapacity)
	// Start the worker
	go sounds.Daemon(witchConfig)

	// New instance of echo
	e := echo.New()

	// Prometheus metrics
	if witchConfig.Metrics {
		prometheus := witchPrometheusMiddleware.NewPrometheus(witchConfig.APIBaseURL)
		prometheus.Use(e)
	}

	// Use logging middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: e.Logger.Output(),
	}))

	// Declare routes
	routes.Routes(e, witchConfig, frontendAssets, apiDocAssets)

	// Start App
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", witchConfig.Port)))
}
