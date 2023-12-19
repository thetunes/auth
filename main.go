package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thetunes/auth/router"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Setup API Router
	router.SetupRoutes(app)

	// Listen on port 3000
	app.Listen(":3000")
}
