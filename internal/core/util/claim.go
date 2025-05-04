package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(payload map[string]interface{}, exp int) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	expiredTime := time.Now().Add(time.Duration(exp) * time.Hour).Unix()
	t.Claims.(jwt.MapClaims)["exp"] = expiredTime

	tokenString, err := t.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (bool, error) {
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
