package main

import "github.com/gofiber/fiber"

func colorHandler(c *fiber.Ctx) {
	for _, light := range hueLights {
		colorParam := c.Params("color")
		// Only change color if in the map
		if _, ok := colorMap[colorParam]; ok {
			bridge.SetLightState(light, getColorState(colorParam))
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
