package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Role     string `gorm:"default:null"`
	Password string `gorm:"default:null"`
	Active   bool   `gorm:"default:true"`
}
