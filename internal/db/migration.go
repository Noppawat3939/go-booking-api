package db

import (
	"go_booking/internal/core/user"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(user.User{})
}
