package repository

import (
	"go_booking/internal/core/domain"

	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db}
}

func (rr *RoomRepository) Create(room *domain.Rooms) error {
	return rr.db.Create(room).Error
}

func (rr *RoomRepository) FindByID(id int) (*domain.Rooms, error) {
	var room *domain.Rooms

	res := rr.db.First(&room, id)

	if res.Error != nil {
		return nil, res.Error
	}

	return room, nil
}

func (rr *RoomRepository) FindAll() ([]*domain.Rooms, error) {
	var rooms []*domain.Rooms
	res := rr.db.Where("active ?", true).Find(&rooms)

	if res.Error != nil {
		return nil, res.Error
	}

	return rooms, nil
}

func (rr *RoomRepository) UpdateByID(id int, room *domain.Rooms) error {
	res := rr.db.Model(&room).Where("id = ?", id).Updates(room)

	return res.Error
}
