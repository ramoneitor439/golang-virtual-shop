package jwtbuilder

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"mystore.com/domain/entities"
	authDtos "mystore.com/dtos/authDtos"
)

func CreateAccessToken(user *entities.User) (*authDtos.SignInResponse, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = user.Id
	claims["sid"] = user.Id
	claims["email"] = user.Email

	roles := ""
	for index, role := range user.Roles {
		if index < len(user.Roles)-1 {
			roles += role.NormalizedName + ","
			continue
		}
		roles += role.NormalizedName
	}

	claims["roles"] = roles
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	secretKey := os.Getenv("jwt_secret_key")
	if secretKey == "" {
		return nil, errors.New("could not load token secret keys")
	}

	tokenString, tokenErr := token.SignedString([]byte(secretKey))
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &authDtos.SignInResponse{
		AccessToken:  tokenString,
		RefreshToken: "",
	}, nil
}
