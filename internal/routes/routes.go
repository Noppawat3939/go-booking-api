package routes

import "github.com/gofiber/fiber/v2"

func InitialRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	HealthRoute(app)
	UserRoutes(api)
	HotelRoutes(api)
}
