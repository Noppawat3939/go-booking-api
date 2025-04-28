package services

import (
	"errors"
	"go_booking/internal/models"
	"go_booking/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(user *models.User) error {
	exit, err := s.repo.FindByUsername(user.Username)

	if err != nil && exit != nil {
		return errors.New("username already exits")
	}

	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}
