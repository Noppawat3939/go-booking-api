package hotels

import "time"

type Hotel struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"default:null"`
	Type        string    `gorm:"default:null"`
	Price       float64   `gorm:"not null"`
	RoomsQty    int       `gorm:"not null"`
	Active      bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedUp   time.Time `gorm:"default:current_timestamp"`
}
