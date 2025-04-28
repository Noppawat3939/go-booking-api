package routes

import (
	"go_booking/internal/controllers"
	"go_booking/internal/db"
	"go_booking/internal/repositories"
	"go_booking/internal/services"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	repo := repositories.NewUserRepository(db.DB)
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	r := router.Group("user")

	r.Get("/", controller.GetUsers)
}
