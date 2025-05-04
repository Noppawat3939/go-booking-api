package service

import (
	"go_booking/internal/core/domain"
	"go_booking/internal/core/port"

	"github.com/gofiber/fiber/v2"
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

func (hs *HotelService) FindOneByID(ctx *fiber.Ctx, id int) (*domain.Hotel, error) {
	return nil, nil
}

func (hs *HotelService) FindAll(ctx *fiber.Ctx) ([]*domain.Hotel, error) {
	return nil, nil
}
