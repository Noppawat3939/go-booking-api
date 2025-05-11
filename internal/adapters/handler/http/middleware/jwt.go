package middleware

import (
	"fmt"
	"go_booking/internal/constants"
	"go_booking/internal/core/domain"
	"go_booking/internal/core/port"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware(ts port.TokenService) fiber.Handler {
	const (
		authType      = "Bearer"
		authHeaderKey = "Authorization"
	)

	return func(c *fiber.Ctx) error {
		authHeader := c.Get(authHeaderKey)

		if authHeader == "" {
			return errorUnAuthorized(c)
		}

		splitted := strings.Split(authHeader, " ")

		if len(splitted) != 2 || splitted[0] != authType {
			return errorUnAuthorized(c)
		}

		token := splitted[1]
		ok, claims, err := ts.VerifyToken(token)

		if err != nil || !ok {
			return errorUnAuthorized(c)
		}

		userID := fmt.Sprintf("%v", claims["id"])
		c.Locals(constants.ContextUserID, userID)

		return c.Next()
	}
}

func errorUnAuthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"success": false,
		"message": domain.UnAuthrorizedMsg,
	})
}
