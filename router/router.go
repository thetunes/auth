package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thetunes/auth/config"
	"github.com/thetunes/auth/handler"
	"github.com/thetunes/auth/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User Group
	login := api.Group("/login")

	// Note: This is just an example, please use a secure secret key
	jwt := middleware.NewAuthMiddleware(config.Config("AUTH_SECRET"))

	// Request token
	login.Post("/auth", handler.Login)

	// Create a protected route
	login.Get("/protected", jwt, handler.Protected)

}
