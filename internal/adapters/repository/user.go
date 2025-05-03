package repository

import (
	"go_booking/internal/core/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(u *domain.User) error {
	return r.db.Create(u).Error
}

func (r *UserRepository) FindByUsername(username string) (*domain.User, error) {
	var user *domain.User

	res := r.db.Where("username = ?", username).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
