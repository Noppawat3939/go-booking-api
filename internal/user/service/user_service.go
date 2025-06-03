package service

import (
	"go_booking/internal/user/model"
	"go_booking/internal/user/port"
	"go_booking/pkg/token"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo  port.UserRepository
	token token.Service
}

func NewUserService(r port.UserRepository, t token.Service) *userService {
	return &userService{repo: r, token: t}
}

func (r *userService) Register(user *model.User) error {
	u := user
	hashedPW, _ := hashPassword(user.Password)
	u.Password = hashedPW

	if err := r.repo.Create(u); err != nil {
		return err
	}

	return nil
}

func (s *userService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)

	if err != nil {
		return "", err
	}

	err = comparePassword(user.Password, password)
	if err != nil {
		return "", err
	}

	t, err := s.token.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return t, nil
}

func (r *userService) GetUserByID(userID string) (*model.User, error) {
	return nil, nil
}

func comparePassword(hashed, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pw))
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", nil
	}

	return string(hashedPassword), nil
}
