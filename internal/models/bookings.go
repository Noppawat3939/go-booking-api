package models

import "time"

type BookingHotel struct {
	ID         uint `gorm:"primaryKey"`
	HotelID    uint
	UserID     uint
	Status     string    `gorm:"not null;check:status in ('pending_payment', 'pending_checkin', 'checked_in', 'checked_out', 'cancelled')"`
	DaysQty    int       `gorm:"not null"`
	CheckInAt  time.Time `gorm:"not null"`
	CheckOutAt time.Time `gorm:"not null"`
	User       User      `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Hotel      Hotel     `gorm:"foreignkey:HotelID;association_foreignkey:ID"`
}
