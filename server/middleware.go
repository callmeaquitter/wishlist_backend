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

func adminMiddleware(c *fiber.Ctx) error {
	specialSession := c.Get("Authorization")
	if specialSession == "" {
		return c.SendString("Unauthorized")
	}

	if specialSession != "user_cnot2oc69lbksn28kko0" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Недостаточно прав для доступа",
		})
	}
	return c.Next()
}

func sellerAuthMiddleware(c *fiber.Ctx) error{
	sellerSessionID := c.Get("seller_Authorization")
	if sellerSessionID == ""{
		return c.SendString("Unauthorized")
	}
	return c.Next()
}