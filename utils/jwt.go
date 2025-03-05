package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(role string) (string, error) {

	JWT_KEY := os.Getenv("JWT_KEY")
	var secretKey = []byte(JWT_KEY)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": role,
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
