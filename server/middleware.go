package server

import (
	"github.com/gofiber/fiber/v2"
)

func authMiddleware(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")

	user, ok := getUserID(authorizationHeader)
	if !ok {
		return c.SendString("Unauthorized")
	}

	c.Locals("user", user)
	return c.Next()
}