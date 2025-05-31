package repository

import (
	"errors"
	"go_booking/internal/user/model"
	"go_booking/internal/user/port"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
	var user model.User

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
