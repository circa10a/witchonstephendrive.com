package routes

import (
	"github.com/amimof/huego"
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/circa10a/witchonstephendrive.com/api" // import generated docs.go
	"github.com/circa10a/witchonstephendrive.com/internal/colors"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// ColorResponse responds with success or failed status string
type ColorResponse struct {
	Status string
}

// :color godoc
// @Summary Change hue lights color
// @Description Change hue lights to supported color defined in color map
// @Produce json
// @Success 200 {object} ColorResponse
// @Failure 400 {object} ColorResponse
// @Router /{color} [post]
// @Param color path string true "Color to change lights to"
func colorHandler(c *fiber.Ctx) error {
	colors := colors.Colors
	color := c.Params("color")
	hueLights := c.Locals("hueLights").([]int)
	bridge := c.Locals("bridge").(*huego.Bridge)
	for _, light := range hueLights {
		// Only change color if in the map
		if _, ok := colors[color]; ok {
			_, err := bridge.SetLightState(light, colors[color])
			if err != nil {
				log.Error(err)
			}
		} else {
			return c.Status(400).JSON(ColorResponse{
				Status: "failed",
			})
		}
	}
	return c.JSON(ColorResponse{
		Status: "success",
	})
}

// Routes instantiates all of the listening context paths
func Routes(app *fiber.App, hueLights []int, bridge *huego.Bridge) {
	// Share huelight id's and bridge connection with route group
	root := app.Group("/", func(c *fiber.Ctx) error {
		c.Locals("hueLights", hueLights)
		c.Locals("bridge", bridge)
		return c.Next()
	})
	// Route to change lights
	root.Post("/:color", colorHandler)
	// Serve frontend static assets
	root.Static("/", "./web")
	// Swagger docs
	root.Static("/api", "./api")
	app.Use("/swagger", swagger.New(swagger.Config{
		URL: "/api/swagger.json",
	}))
}
