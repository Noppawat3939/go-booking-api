package domain

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"default:null"`
	Address     string `gorm:"not null"`
	City        string `gorm:"not null"`
	State       string `gorm:"not null"`
	Country     string `gorm:"not null"`
	PostCode    string `gorm:"not null"`
	PhoneNumber string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Website     string `gorm:"not null"`
	Active      bool   `gorm:"default:true"`
}

type HotelFacility struct {
	ID        uint `gorm:"primaryKey"`
	HotelID   int
	Wifi      bool  `gorm:"default:false"`
	BreakFast bool  `gorm:"default:false"`
	Hotel     Hotel `gorm:"foreignkey:HotelID;association_foreignkey:ID"`
}
