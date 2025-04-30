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

func (ur *UserRepository) CreateUser(u *domain.User) error {
	return ur.db.Create(u).Error
}
