package repository

import (
	"go_booking/internal/core/domain"

	"gorm.io/gorm"
)

type HotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) *HotelRepository {
	return &HotelRepository{db}
}

func (hr *HotelRepository) Create(hotel *domain.Hotel) error {
	return hr.db.Create(hotel).Error
}

func (hr *HotelRepository) FindByID(id int) (*domain.Hotel, error) {
	return nil, nil
}

func (hr *HotelRepository) FindAll() ([]*domain.Hotel, error) {
	return nil, nil
}
