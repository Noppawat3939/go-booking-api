package port

import (
	"go_booking/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type TokenService interface {
	CreateToken(payload *domain.User) (string, error)
	VerifyToken(tokenString string) (bool, error)
}

type AuthService interface {
	// return a token
	Login(ctx *fiber.Ctx, username, password string) (string, error)
}
