package routes

import (
	"fmt"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/internal/colors"
	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

func colorHandler(c *fiber.Ctx) {
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

// Routes instantiates all of the listening context paths
func Routes(app *fiber.App, hueLights []int, bridge *huego.Bridge) {
	// Share huelight id's and bridge connection with route group
	root := app.Group("/", func(c *fiber.Ctx) {
		c.Locals("hueLights", hueLights)
		c.Locals("bridge", bridge)
		c.Next()
	})
	// // Route to change lights
	root.Post("/:color", colorHandler)
	// Serve frontend static assets
	root.Static("/", "./web")
}
