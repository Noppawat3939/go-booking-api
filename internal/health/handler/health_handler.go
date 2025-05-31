package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HealthHander(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sqlDB, err := db.DB()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "fail",
				"message": "failed to get database connection",
			})
		}

		if err := sqlDB.Ping(); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"status": "fail",
				"db":     "unhealthy",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":       "ok",
			"db":           "healthy",
			"requested_at": time.Now(),
		})
	}
}
