package routes

import (
	"go_booking/internal/db"
	"go_booking/internal/handlers"
	"go_booking/internal/repositories"
	"go_booking/internal/services"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	repo := repositories.NewUserRepository(db.DB)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHanlder(service)

	r := router.Group("user")

	r.Get("/", handler.GetUsers)
	r.Post("/", handler.CreateUser)
}
