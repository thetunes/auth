package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/thetunes/auth/config"
	"github.com/thetunes/auth/router"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Setup API Router
	router.SetupRoutes(app)

	// Allow our domain to access
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     strings.Join(config.GetAllowedOrigins(), ","),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// Listen on port 3000
	app.Listen(":3000")
}
