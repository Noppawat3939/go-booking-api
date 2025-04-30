package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRouter(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	UserRoutes(v1, db)
}
