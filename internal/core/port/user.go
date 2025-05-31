package port

import (
	"go_booking/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
	FindOneByID(id int) (*domain.User, error)
}

type UserService interface {
	Register(ctx *fiber.Ctx, user *domain.User) error
	GetUser(id int) (*domain.User, error)
}
