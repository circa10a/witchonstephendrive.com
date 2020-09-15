package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/amimof/huego"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/circa10a/witchonstephendrive.com/internal/routes"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"
)

var (
	port      *int
	metrics   *bool
	hueUser   *string
	bridge    *huego.Bridge
	hueLights []int
)

func init() {
	// Parse flags
	flags()
	// Find hue bridge ip
	hueBridge, err := huego.Discover()
	if err != nil {
		log.Fatal(err)
	}
	// Assign bridge location to global var
	bridge = hueBridge
	// Authenticate against bridge api
	bridge.Login(*hueUser)
}

// Read flags from command line args and set defaults
func flags() {
	// Args
	port = flag.Int("port", 8080, "Listening port")
	metrics = flag.Bool("metrics", true, "Enable prometheus metrics")
	hueUser = flag.String("hue-user", os.Getenv("HUE_USER"), "ID to connect to hue bridge")
	lightsStr := flag.String("hue-lights", os.Getenv("HUE_LIGHTS"), "Light ID's to change")
	flag.Parse()

	// Parse string input to slice of ints
	hueLights = utils.StrToIntSlice(strings.Fields(*lightsStr))

	// Validation
	if *hueUser == "" {
		log.Fatal("HUE_USER not set")
	}
	if len(hueLights) == 0 {
		log.Fatal("HUE_LIGHTS not set")
	}
}

// @title witchonstephendrive.com
// @version 0.1.0
// @description Control my lights for Halloween
// @contact.name Caleb Lemoine
// @contact.email caleblemoine@gmail.com
// @license.name MIT
// @license.url https://raw.githubusercontent.com/circa10a/witchonstephendrive.com/master/LICENSE
// @host witchonstephendrive.com
// @BasePath /
// @Schemes https
func main() {
	// New instance of fiber
	app := fiber.New()
	// Prometheus metrics
	if *metrics {
		prometheus := fiberprometheus.New("witch-metrics")
		prometheus.RegisterAt(app, "/metrics")
		app.Use(prometheus.Middleware)
	}
	// Use logging middleware
	app.Use(logger.New())
	// Declare routes
	routes.Routes(app, hueLights, bridge)
	// Start App
	app.Listen(fmt.Sprintf(":%v", *port))
}
