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

func (hr *HotelRepository) FindAll() ([]*domain.Hotel, error) {
	var hotels []*domain.Hotel
	res := hr.db.Where("active = ?", true).Find(&hotels)

	if res.Error != nil {
		return nil, res.Error
	}

	return hotels, nil
}

func (hr *HotelRepository) FindByID(id int) (*domain.Hotel, error) {
	var hotel *domain.Hotel
	res := hr.db.First(&hotel, id)

	if res.Error != nil {
		return nil, res.Error
	}

	return hotel, nil
}

func (hr *HotelRepository) UpdateByID(id int, hotel *domain.Hotel) error {
	res := hr.db.Model(&hotel).Where("id = ?", id).Updates(hotel)

	return res.Error
}
