package db

import (
	"go_booking/internal/core/domain"
	"go_booking/internal/user/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{}, domain.Hotel{}, domain.Rooms{}, domain.RoomFacilities{})
}
