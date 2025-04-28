package repositories

import (
	"go_booking/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, username).Error
	return &user, err
}

func (r *UserRepository) FindOne(username string, role string) (*models.User, error) {
	var user models.User

	err := r.db.Where("username = ? AND role = ?", username, role).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}
