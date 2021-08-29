package main

import (
	"embed"
	"fmt"

	"github.com/circa10a/witchonstephendrive.com/controllers/sounds"
	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/routes"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// all environment variables for configuration expect WITCH_ prefix
const envVarPrefix = "witch"

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
	witchConfig := &config.WitchConfig{}
	err := envconfig.Process(envVarPrefix, witchConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Show HAPPY HALLOWEEN banner
	if witchConfig.ShowBanner {
		witchConfig.PrintBanner()
	}

	// Logger Config
	err = witchConfig.InitLogger()
	if err != nil {
		log.Fatal(err)
	}

	// Hue Lights
	// Start scheduler to regularly redescover bridge IP in the event DHCP changes it
	witchConfig.InitHue()
	// Start scheduler to set default light colors (if enabled)
	witchConfig.InitDefaultColorsScheduler()
	// Start schedulers to turn lights on/off
	witchConfig.InitHueLightsScheduler()

	// Sounds
	// Google Assistant Relay Config such as endpoint and client
	witchConfig.InitAssistantRelayConfig()
	// Creates initial capped sounds queue
	witchConfig.InitSoundQueue()
	// Start the sound queue worker
	go sounds.InitDaemon(witchConfig)

	// Rest API
	// Configure echo server
	e := witchConfig.InitEchoConfig(frontendAssets, apiDocAssets)
	// Declare routes + handlers
	routes.Routes(e, witchConfig, frontendAssets, apiDocAssets)
	// Start Listener
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", witchConfig.Port)))
}
