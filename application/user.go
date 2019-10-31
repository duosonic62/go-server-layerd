package application

import "sample-sever/domain"

func GetUsers() (*[]domain.User, error) {
	users, error := domain.GetUsers()
	if error != nil {
		return nil, error
	}
	return users, error
}
