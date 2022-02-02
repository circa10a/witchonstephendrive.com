package main

import (
	"embed"
	"fmt"

	"github.com/circa10a/witchonstephendrive.com/controllers/sounds"
	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/routes"
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
	witchConfig := config.New()
	// Start the sound queue worker
	go sounds.InitDaemon(witchConfig)

	// Rest API
	// Configure echo server
	e := witchConfig.InitEchoConfig(frontendAssets, apiDocAssets)
	// Declare routes + handlers
	routes.Routes(e, witchConfig, frontendAssets, apiDocAssets)
	// Start Listener
	log.Fatal(e.Start(fmt.Sprintf(":%d", witchConfig.Port)))
}
