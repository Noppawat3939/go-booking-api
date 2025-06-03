package port

import "go_booking/internal/user/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByUsername(username string) (*model.User, error)
	FindByID(id string) (*model.User, error)
}

type UserService interface {
	Register(user *model.User) error
	Login(username, password string) (string, error)
	GetUserByID(id string) (*model.UserResponse, error)
}
