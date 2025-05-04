package dto

import (
	"time"
)

type Hotel struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	PostCode    string    `json:"post_code"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Website     string    `json:"website"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type CreateHotelRequest struct {
	Hotel
	ID        int       `json:"-"` //omit id
	CreatedAt time.Time `json:"-"` // omit created_at
	UpdatedAt time.Time `json:"-"` // omit updated_at
	DeletedAt time.Time `json:"-"` // omit deleted_at
}

type CreateHotelResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	PostCode    string `json:"post_code"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Website     string `json:"website"`
}
