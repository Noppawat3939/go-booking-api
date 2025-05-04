package db

import (
	"go_booking/internal/core/domain"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(domain.User{}, domain.Hotel{})
}
