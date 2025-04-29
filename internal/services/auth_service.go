package services

import (
	"go_booking/internal/models"
	"go_booking/internal/repositories"
)

type AuthService struct {
	repo *repositories.UserRepository
}

func NewAuthService(r *repositories.UserRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *UserService) Login(username string, password string) error {
	return nil
}

func (s *UserService) Register(user *models.User) error {
	return nil
}
