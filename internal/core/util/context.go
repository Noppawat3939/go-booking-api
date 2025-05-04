package util

import (
	"go_booking/internal/constants"

	"github.com/gofiber/fiber/v2"
)

func GetUserIDFromContext(c *fiber.Ctx) (string, bool) {
	userID, ok := c.Locals(constants.ContextUserID).(string)

	return userID, ok
}
