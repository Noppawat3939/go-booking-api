package service

import (
	"go_booking/internal/core/domain"
	"go_booking/internal/core/port"
	"go_booking/internal/core/util"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	repo port.UserRepositoyry
}

func NewUserService(repo port.UserRepositoyry) *UserService {
	return &UserService{repo}
}

func (us *UserService) Register(c *fiber.Ctx, user *domain.User) error {
	hashedPassword, err := util.HashPassword(user.Password)

	if err != nil {
		return domain.ErrInternal
	}

	user.Password = hashedPassword

	if err := us.repo.CreateUser(user); err != nil {
		return domain.ErrConflict
	}

	return nil
}
