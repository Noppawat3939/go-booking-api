package domain

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	HotelID  uint
	Name     string
	Type     string
	Price    float64
	Discount float64
	Active   bool
	Hotel    Hotel `gorm:"foreignkey:HotelID;association_foreignkey:ID"`
}

type RoomFacilities struct {
	ID     uint `gorm:"primaryKey"`
	RoomID uint
	Values string
	Rooms  Rooms `gorm:"foreignkey:RoomID;association_foreignkey:ID"`
}
