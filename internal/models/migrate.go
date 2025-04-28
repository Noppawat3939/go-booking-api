package models

import "gorm.io/gorm"

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Hotel{}, &HotelFacility{}, &BookingHotel{})
}
