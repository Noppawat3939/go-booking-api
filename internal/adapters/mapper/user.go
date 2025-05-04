package mapper

import (
	"go_booking/internal/adapters/handler/dto"
	"go_booking/internal/core/domain"
)

func ToCreateUserReq(r dto.CreateUserRequest) domain.User {
	return domain.User{
		Username: r.Username,
		Password: r.Password,
		Role:     r.Role,
	}
}

func ToCreateUserRes(u *domain.User) dto.CreateUserResponse {
	return dto.CreateUserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
