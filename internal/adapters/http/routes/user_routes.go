package routes

import (
	"go_booking/internal/adapters/http/handler"
	"go_booking/internal/adapters/repository"
	"go_booking/internal/core/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(router fiber.Router, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	service := user.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	r := router.Group("/users")

	r.Post("/", handler.Register)
}
