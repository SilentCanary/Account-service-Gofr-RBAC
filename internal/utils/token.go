package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("SECRET") // 🔒 later use env var

func GenerateJWT(userID string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"roles":  roles,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
