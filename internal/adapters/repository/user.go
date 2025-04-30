package repository

import (
	"go_booking/internal/core/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(u *user.User) error {
	return r.db.Create(u).Error
}
