package port

import (
	"go_booking/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type UserRepositoyry interface {
	CreateUser(user *domain.User) error
}

type UserService interface {
	Register(ctx *fiber.Ctx, user *domain.User) error
}
