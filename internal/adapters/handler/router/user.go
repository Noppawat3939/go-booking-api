package router

import (
	"go_booking/internal/adapters/handler/http"
	"go_booking/internal/adapters/repository"
	"go_booking/internal/core/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRouter(router fiber.Router, db *gorm.DB) {
	// initialized User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	// initialized Auth
	tokenService := service.NewTokenService()
	authService := service.NewAuthService(userRepo, tokenService)
	authHandler := http.NewAuthHandler(authService)

	r := router.Group("/user")

	r.Post("/", userHandler.Register)
	r.Post("/login", authHandler.Login)
}
