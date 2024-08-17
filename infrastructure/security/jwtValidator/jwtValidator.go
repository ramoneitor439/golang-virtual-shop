package jwtvalidator

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func Validate(accessToken string) error {
	secretKey := os.Getenv("jwt_secret_key")
	token, _ := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error obtaining token information")
		}
		return secretKey, nil
	})

	if token == nil {
		return errors.New("invalid token")
	}

	return nil
}
