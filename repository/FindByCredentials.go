package repository

import (
	"errors"

	"github.com/thetunes/auth/model"
)

func FindByCredentials(username, password string) (*model.User, error) {
	// Here you would query your database for the user with the given email
	if username == "test@mail.com" && password == "test12345" {
		return &model.User{
			ID:             1,
			Username:       "test@mail.com",
			Password:       "test12345",
			FavoritePhrase: "Hello, World!",
		}, nil
	}
	return nil, errors.New("user not found")
}
