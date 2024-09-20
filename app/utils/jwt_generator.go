package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken() (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
