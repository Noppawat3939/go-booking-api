package port

import (
	"go_booking/internal/core/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	CreateToken(payload *domain.User) (string, error)
	VerifyToken(tokenString string) (bool, jwt.MapClaims, error)
}

type AuthService interface {
	Login(ctx *fiber.Ctx, username, password string) (string, error)
}
