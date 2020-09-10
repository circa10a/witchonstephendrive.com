package main

import (
	"flag"
	"os"
	"strings"

	"github.com/amimof/huego"
	"github.com/ansrivas/fiberprometheus"
	"github.com/circa10a/witchonstephendrive.com/internal/routes"
	"github.com/circa10a/witchonstephendrive.com/pkg/utils"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
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
	app.Use(middleware.Logger())
	// Declare routes
	routes.Routes(app, hueLights, bridge)
	// Start App
	app.Listen(*port)
}
