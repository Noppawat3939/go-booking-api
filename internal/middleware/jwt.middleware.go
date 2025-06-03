package middleware

import (
	"fmt"
	"go_booking/pkg/token"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JwtBearerMiddleware(t token.Service) fiber.Handler {
	const (
		authType   = "Bearer"
		headersKey = "Authorization"
	)

	return func(c *fiber.Ctx) error {
		header := c.Get(headersKey)

		if header == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "unauthroized"})
		}

		s := strings.Split(header, " ")

		if len(s) != 2 || s[0] != authType {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "unauthroized"})
		}

		tk := s[1]
		ok, claims, err := t.ValidateToken(tk)

		if err != nil || !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "unauthroized"})
		}

		userID := fmt.Sprintf("%v", claims["user_id"])
		c.Locals("user_id", userID)

		return c.Next()
	}
}
