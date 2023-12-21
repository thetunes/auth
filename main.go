package main

import (
	"github.com/gofiber/fiber/v2"
  	"github.com/thetunes/auth/router"
        "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Setup API Router
	router.SetupRoutes(app)

        app.Use(cors.New(cors.Config{
            AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
            AllowOrigins:     "*",
            AllowCredentials: true,
            AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
        }))

	// Listen on port 3000
	app.Listen(":80")
}
