package service

import (
	"go_booking/internal/core/domain"
	"go_booking/internal/core/util"
)

type JWTService struct{}

func NewTokenService() *JWTService {
	return &JWTService{}
}

func (s *JWTService) CreateToken(user *domain.User) (string, error) {
	tokenPayload := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	}

	return util.CreateToken(tokenPayload, 3)
}
