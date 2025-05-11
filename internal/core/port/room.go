package port

import "go_booking/internal/core/domain"

type RoomRepository interface {
	Create(room *domain.Rooms) error
	FindByID(id int) (*domain.Rooms, error)
	FindAll() ([]*domain.Rooms, error)
	UpdateByID(id int, room *domain.Rooms) error
}

type RoomService interface {
	CreateRoom(room *domain.Rooms) error
	FindOneByID(id int) (*domain.Rooms, error)
	FindAll() ([]*domain.Rooms, error)
	UpdateByID(id int, room *domain.Rooms) error
}
