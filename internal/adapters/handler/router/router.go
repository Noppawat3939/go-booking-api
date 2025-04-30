package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitialRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	UserRouter(v1, db)
}
