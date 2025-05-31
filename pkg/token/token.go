package token

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	UserID string `json:"user_id"`
}

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (bool, jwt.MapClaims, error)
}
