package main

import (
	"embed"
	"fmt"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/internal/routes"
	"github.com/go-resty/resty/v2"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
// @license.name MIT
// @license.url https://raw.githubusercontent.com/circa10a/witchonstephendrive.com/main/LICENSE
// @host witchonstephendrive.com
// @BasePath /
// @Schemes https
func main() {
	// Setup config
	var witchConfig config.WitchConfig
	err := envconfig.Process("witch", &witchConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Hue config
	// Find hue bridge ip
	hueBridge, err := huego.Discover()
	if err != nil {
		log.Fatal(err)
	}
	witchConfig.Bridge = hueBridge
	// Authenticate against bridge api
	witchConfig.Bridge.Login(witchConfig.HueUser)
	// Store all light data to be used later
	for _, lightID := range witchConfig.HueLights {
		light, err := witchConfig.Bridge.GetLight(lightID)
		if err != nil {
			log.Fatal(err)
		}
		witchConfig.HueLightsStructs = append(witchConfig.HueLightsStructs, *light)
	}

	// Create client to be used with assistant relay
	assitantRelatEndpoint := fmt.Sprintf("%v:%v", witchConfig.AssistantRelayHost, witchConfig.AssistantRelayPort)
	witchConfig.Client = resty.New().SetHostURL(assitantRelatEndpoint).SetHeader("Content-Type", "application/json")

	// New instance of echo
	e := echo.New()
	// Prometheus metrics
	if witchConfig.Metrics {
		prometheus.NewPrometheus("echo", nil).Use(e)
	}
	// Use logging middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: e.Logger.Output(),
	}))
	// Declare routes
	routes.Routes(e, witchConfig, frontendAssets, apiDocAssets)
	// Start App
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", witchConfig.Port)))
}
