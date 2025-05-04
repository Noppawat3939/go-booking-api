package port

import (
	"go_booking/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type HotelRepository interface {
	Create(hotel *domain.Hotel) error
	FindByID(id int) (*domain.Hotel, error)
	FindAll() ([]*domain.Hotel, error)
}

type HotelService interface {
	CreateHotel(hotel *domain.Hotel) error
	FindOneByID(id int) (*domain.Hotel, error)
	FindAll(ctx *fiber.Ctx) ([]*domain.Hotel, error)
}
