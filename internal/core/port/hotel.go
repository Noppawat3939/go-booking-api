package port

import (
	"go_booking/internal/core/domain"
)

type HotelRepository interface {
	Create(hotel *domain.Hotel) error
	FindByID(id int) (*domain.Hotel, error)
	FindAll() ([]*domain.Hotel, error)
}

type HotelService interface {
	CreateHotel(hotel *domain.Hotel) error
	FindOneByID(id int) (*domain.Hotel, error)
	FindAll() ([]*domain.Hotel, error)
}
