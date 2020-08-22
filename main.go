package main

import (
	"flag"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/amimof/huego"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

var (
	port      *int
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
	hueUser = flag.String("hue-user", os.Getenv("HUE_USER"), "ID to connect to hue bridge")
	lightsStr := flag.String("lights", os.Getenv("HUE_LIGHTS"), "Light ID's to change")
	flag.Parse()

	// Ensure converstion
	hueLights = strToIntSlice(strings.Fields(*lightsStr))

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
	// Use logging middleware
	app.Use(middleware.Logger())
	// Declare routes
	routes(app)
	// Start App
	app.Listen(*port)
}
