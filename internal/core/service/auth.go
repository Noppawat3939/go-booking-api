package service

import (
	"go_booking/internal/core/port"
	"go_booking/internal/core/util"

	"github.com/gofiber/fiber/v2"
)

type AuthService struct {
	repo port.UserRepository
	ts   port.TokenService
}

func NewAuthService(repo port.UserRepository, ts port.TokenService) *AuthService {
	return &AuthService{repo, ts}
}

func (as *AuthService) Login(c *fiber.Ctx, username, password string) (string, error) {
	user, err := as.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if err = util.ComparePassword(password, user.Password); err != nil {
		return "", err
	}

	token, err := as.ts.CreateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
