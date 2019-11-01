package application

import (
	"sample-sever/domain"
	"sample-sever/domain/errors"
)

func GetUsers() (*[]domain.User, error) {
	users, error := domain.GetUsers()
	if error != nil {
		appError := errors.New(500, "DB ERROR.", "DB ERROR.", error)
		return nil, appError
	}
	return users, error
}
