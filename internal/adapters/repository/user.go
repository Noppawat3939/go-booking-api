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

func (ur *UserRepository) FindByUsername(username string) (*domain.User, error) {
	var user *domain.User

	res := ur.db.Where("username = ?", username).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (ur *UserRepository) FindOneByID(id string) (*domain.User, error) {
	var user *domain.User
	res := ur.db.Where("id = ?", id).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil

}
