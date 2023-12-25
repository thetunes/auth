// handler.go
package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/thetunes/auth/config"
	"github.com/thetunes/auth/model"
	"github.com/thetunes/auth/repository"
)

// Login route
func Login(c *fiber.Ctx) error {
	loginRequest := new(model.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := repository.FindByCredentials(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	day := time.Hour * 24

	claims := jwt.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(day * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Config("AUTH_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(model.LoginResponse{
		Token: t,
	})
}

func Protected(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	// Assuming you have a "fav" claim in your JWT token
	favPhrase := claims["fav"].(string)
	return c.SendString("Welcome 👋" + username + " " + favPhrase)
}
