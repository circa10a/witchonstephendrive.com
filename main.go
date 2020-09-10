package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/amimof/huego"
	"github.com/ansrivas/fiberprometheus"
	"github.com/circa10a/witchonstephendrive.com/internal/colors"
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

func colorHandler(c *fiber.Ctx) {
	colors := colors.Colors
	color := c.Params("color")
	for _, light := range hueLights {
		// Only change color if in the map
		if _, ok := colors[color]; ok {
			_, err := bridge.SetLightState(light, colors[color])
			if err != nil {
				log.Error(err)
			}
			c.JSON(fiber.Map{
				"status": fmt.Sprintf("set to %s", color),
			})
		} else {
			c.JSON(fiber.Map{
				"status": fmt.Sprintf("%s not found", color),
			})
		}
	}
}

func routes(app *fiber.App) {
	root := app.Group("/")
	// // Route to change lights
	root.Post("/:color", colorHandler)
	// Serve frontend static assets
	root.Static("/", "./web")
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
	routes(app)
	// Start App
	app.Listen(*port)
}
