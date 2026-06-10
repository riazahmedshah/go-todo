package main

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateToken(userId int64) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
