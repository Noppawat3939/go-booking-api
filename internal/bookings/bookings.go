package bookings

import "time"

type BookingHotel struct {
	ID         uint      `gorm:"primaryKey"`
	HotelID    uint      `gorm:"not null;index"`
	UserID     uint      `gorm:"not null;index"`
	Status     string    `gorm:"not null;check:status in ('pending_payment', 'pending_checkin', 'checked_in', 'checked_out', 'cancelled')"`
	DaysQty    int       `gorm:"not null"`
	CheckInAt  time.Time `gorm:"not null"`
	CheckOutAt time.Time `gorm:"not null"`
}
