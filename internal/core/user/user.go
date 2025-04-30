package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:null"`
	Active   bool   `gorm:"default:true"`
}

type Repository interface {
	Create(user *User) error
}
