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
	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)
	uh := http.NewUserHandler(us)

	// initialized Auth
	ts := service.NewTokenService()
	as := service.NewAuthService(ur, ts)
	ah := http.NewAuthHandler(as)

	r := router.Group("/user")

	r.Post("/", uh.Register)
	r.Post("/login", ah.Login)
}
