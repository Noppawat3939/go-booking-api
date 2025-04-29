package models

import (
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string  `gorm:"not null"`
	Description string  `gorm:"default:null"`
	Type        string  `gorm:"default:null"`
	Price       float64 `gorm:"not null"`
	RoomsQty    int     `gorm:"not null"`
	Active      bool    `gorm:"default:true"`
}

type HotelFacility struct {
	ID        uint `gorm:"primaryKey"`
	HotelID   int
	Wifi      bool  `gorm:"default:false"`
	BreakFast bool  `gorm:"default:false"`
	Hotel     Hotel `gorm:"foreignkey:HotelID;association_foreignkey:ID"`
}
