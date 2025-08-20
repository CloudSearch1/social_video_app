package main

import (
	"log"

	"douyin_app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	setupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

func setupRoutes(app *fiber.App) {
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// User routes
	routes.UserRoutes(app)
	// Video routes
	routes.VideoRoutes(app)
}