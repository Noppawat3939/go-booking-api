package user

import (
	"go_booking/internal/middleware"
	"go_booking/internal/user/handler"
	"go_booking/internal/user/repository"
	"go_booking/internal/user/service"
	"go_booking/pkg/token"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router fiber.Router, db *gorm.DB) {

	repo := repository.NewUserRepository(db)
	tokenSvc := token.NewTokenService()
	service := service.NewUserService(repo, tokenSvc)
	handler := handler.NewUserHandler(service)

	r := router.Group("user")

	r.Post("/", middleware.JwtBearerMiddleware(tokenSvc), handler.GetUser)
	r.Post("/register", handler.Register)
	r.Post("/login", handler.Login)
}
