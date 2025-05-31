package routes

import (
	"go_booking/internal/health/handler"
	"go_booking/internal/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	api.Get("/healthz", handler.HealthHander(db))

	user.RegisterUserRoutes(api, db)
}
