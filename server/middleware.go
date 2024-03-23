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

	fmt.Println(specialSession)
	specialSession = strings.TrimPrefix(specialSession, "Bearer ")
	session, ok := db.FindSession(specialSession)
	if !ok {
		return c.SendString("Unauthorized")
	}

	if specialSession != "session_cnvdk9k69lbm5c1vej1g" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Недостаточно прав для доступа",
		})
	}

	fmt.Println(specialSession)
	fmt.Println(session)
	c.Locals("user", session.ID)
	return c.Next()
}

func sellerAuthMiddleware(c *fiber.Ctx) error {
	sellerSessionID := c.Get("seller_Authorization")
	if sellerSessionID == "" {
		return c.SendString("Unauthorized")
	}

	sellerSessionID = strings.TrimPrefix(sellerSessionID, "Bearer ")
	sellerSession, ok := db.FindSellerSession(sellerSessionID)
	if !ok {
		return c.SendString("Unauthorized")
	}

	fmt.Println(sellerSessionID)
	fmt.Println(sellerSession)
	c.Locals("seller", sellerSession.SellerID)
	return c.Next()
}
