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

//go:embed web
var frontendAssets embed.FS

//go:embed api
var apiDocAssets embed.FS

func init() {
	// Logger Config
	config.InitLogger()
}

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
	err := envconfig.Process("", witchConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Hue Lights
	// Goroutine to regularly redescovers bridge IP in the event DHCP changes it
	go witchConfig.InitHue()

	// Sounds
	// Google Assistant Relay Config such as endpoint and client
	witchConfig.InitAssistantRelayConfig()
	// Creates initial capped sounds queue
	witchConfig.InitSoundQueue()
	// Start the sound queue worker
	go sounds.Daemon(witchConfig)

	// Rest API
	// Configure echo server
	e := witchConfig.InitEchoConfig(frontendAssets, apiDocAssets)
	// Declare routes + handlers
	routes.Routes(e, witchConfig, frontendAssets, apiDocAssets)
	// Start Listener
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", witchConfig.Port)))
}
