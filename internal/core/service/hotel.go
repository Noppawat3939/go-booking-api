package service

import (
	"go_booking/internal/core/domain"
	"go_booking/internal/core/port"
)

type HotelService struct {
	repo port.HotelRepository
}

func NewHotelService(repo port.HotelRepository) *HotelService {
	return &HotelService{repo}
}

func (hs *HotelService) CreateHotel(hotel *domain.Hotel) error {
	if err := hs.repo.Create(hotel); err != nil {
		return err
	}

	return nil
}

func (hs *HotelService) FindAll(skip, limit int) ([]*domain.Hotel, error) {
	hotels, err := hs.repo.FindAll(skip, limit)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func (hs *HotelService) FindOneByID(id int) (*domain.Hotel, error) {
	hotel, err := hs.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return hotel, nil
}

func (hs *HotelService) UpdateByID(id int, hotel *domain.Hotel) error {
	err := hs.repo.UpdateByID(id, hotel)

	if err != nil {
		return err
	}

	return nil
}
