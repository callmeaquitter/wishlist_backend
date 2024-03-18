package server

import (
	"fmt"
	"strings"
	"wishlist/db"

	"github.com/gofiber/fiber/v2"
)

func authMiddleware(c *fiber.Ctx) error {
	sessionID := c.Get("Authorization")
	if sessionID == "" {
		return c.SendString("Unauthorized")
	}

	sessionID = strings.TrimPrefix(sessionID, "Bearer ")
	session, ok := db.FindSession(sessionID)
	if !ok {
		return c.SendString("Unauthorized")
	}
	
	fmt.Println(sessionID)
	fmt.Println(session)
	c.Locals("user", session.UserID)
	return c.Next()
}
