package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/circa10a/witchonstephendrive.com/controllers/sounds"
	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/internal/routes"
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
	w := config.New()

	// Configure echo server
	e := w.InitEchoConfig(frontendAssets, apiDocAssets)

	ctx, cancel := context.WithCancel(context.Background())
	// Ensure we can cancel all of our goroutines
	go func(ctx context.Context) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c

		cancel()
		if err := e.Shutdown(ctx); err != nil {
			if err != context.Canceled {
				log.Fatal(err, "Error shutting down web server")
			}
		}
	}(ctx)

	// Hue Lights
	// Start scheduler to regularly redescover bridge IP in the event DHCP changes it
	go w.InitHue(ctx)
	// Start scheduler to set default light colors (if enabled)
	w.InitDefaultColorsScheduler()
	// Start schedulers to turn lights on/off
	w.InitHueLightsScheduler()

	// Sounds
	// Start the sound queue worker
	go sounds.InitDaemon(ctx, w)

	// REST API
	// Declare routes + handlers
	routes.Routes(e, w, frontendAssets, apiDocAssets)
	// Start Listener
	if err := e.Start(fmt.Sprintf(":%d", w.Port)); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err, "Error starting web server")
		}
		log.Info("Web server shutdown successfully")
	}
}
