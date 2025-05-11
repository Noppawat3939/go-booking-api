package util

import (
	"go_booking/internal/constants"

	"github.com/gofiber/fiber/v2"
)

func GetUserIDFromContext(c *fiber.Ctx) (string, bool) {
	userID, ok := c.Locals(constants.ContextUserID).(string)

	return userID, ok
}

func GetPaginationParams(c *fiber.Ctx) (page, limit, offset int) {
	page = c.QueryInt("page", 1)
	limit = c.QueryInt("limit", 100)

	if page < 1 {
		page = 1
	}

	offset = (page - 1) * limit
	return
}
