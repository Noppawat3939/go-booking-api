package dto

type CreateHotelRequest struct {
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
