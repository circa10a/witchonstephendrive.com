package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

func colorHandler(c *fiber.Ctx) {
	color := c.Params("color")
	for _, light := range hueLights {
		// Only change color if in the mgpriap
		if _, ok := colors[color]; ok {
			_, err := bridge.SetLightState(light, colors.getColorState(color))
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

func healthHandler(c *fiber.Ctx) {
	c.JSON(fiber.Map{
		"status": "ok",
	})
}

func routes(app *fiber.App) {
	root := app.Group("/")
	// Route to change lights
	root.Post("/:color", colorHandler)
	// Serve frontend static assets
	root.Static("/", "./public")
	// Health check
	root.Get("/health", healthHandler)
}
