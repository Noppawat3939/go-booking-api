package service

import (
	"go_booking/internal/core/domain"
	"go_booking/internal/core/port"
)

type RoomService struct {
	repo port.RoomRepository
}

func NewRoomService(repo port.RoomRepository) *RoomService {
	return &RoomService{repo}
}

func (rs *RoomService) CreateRoom(room *domain.Rooms) error {
	if err := rs.repo.Create(room); err != nil {
		return err
	}

	return nil
}

func (rs *RoomService) FindOneByID(id int) (*domain.Rooms, error) {
	room, err := rs.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (rs *RoomService) FindAll() ([]*domain.Rooms, error) {
	rooms, err := rs.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (rs *RoomService) UpdateByID(id int, room *domain.Rooms) error {
	if err := rs.repo.UpdateByID(id, room); err != nil {
		return err
	}

	return nil
}
