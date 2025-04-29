package routes

import (
	"go_booking/internal/db"
	"go_booking/internal/handlers"
	"go_booking/internal/repositories"
	"go_booking/internal/services"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	repo := repositories.NewUserRepository(db.DB)
	service := services.NewAuthService(repo)
	handler := handlers.NewAuthHandler(service)

	r := router.Group("auth")

	r.Post("/login", handler.LoginUser)
	r.Post("/register", handler.RegisterUser)
}
