package main

import (
	"embed"
	"fmt"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/internal/routes"
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

type witchConfig struct {
	Port      int           `envconfig:"PORT" default:"8080"`
	Metrics   bool          `envconfig:"METRICS" default:"true"`
	HueUser   string        `envconfig:"HUE_USER" required:"true"`
	HueLights []int         `envconfig:"HUE_LIGHTS" required:"true" split_words:"true"`
	Bridge    *huego.Bridge `ignored:"true"`
}

// @title witchonstephendrive.com
// @version 0.1.0
// @description Control my lights for Halloween
// @contact.name Caleb Lemoine
// @contact.email caleblemoine@gmail.com
// @license.name MIT
// @license.url https://raw.githubusercontent.com/circa10a/witchonstephendrive.com/main/LICENSE
// @host witchonstephendrive.com
// @BasePath /
// @Schemes https
func main() {
	// setup config
	var witchConfig witchConfig
	err := envconfig.Process("witch", &witchConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Find hue bridge ip
	hueBridge, err := huego.Discover()
	if err != nil {
		log.Fatal(err)
	}
	// Assign bridge location to global var
	witchConfig.Bridge = hueBridge
	// Authenticate against bridge api
	witchConfig.Bridge.Login(witchConfig.HueUser)

	// New instance of echo
	e := echo.New()
	// Prometheus metrics
	if witchConfig.Metrics {
		p := prometheus.NewPrometheus("echo", nil)
		p.Use(e)
	}
	// Use logging middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: e.Logger.Output(),
	}))
	// Declare routes
	routes.Routes(e, witchConfig.HueLights, witchConfig.Bridge, frontendAssets, apiDocAssets)
	// Start App
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", witchConfig.Port)))
}
