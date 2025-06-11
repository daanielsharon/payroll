package utils

import (
	"shared/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userID, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(config.LoadConfig().JWTSecret))
}
