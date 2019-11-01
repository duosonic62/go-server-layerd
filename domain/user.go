package domain

import "sample-sever/infra"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUsers() (*[]User, error) {
	dtos, error := infra.GetUserDtos()
	if error != nil {
		return nil, error
	}
	users := make([]User, len(dtos))

	// Convert: Dto -> Domain
	for _, dto := range dtos {
		user := User{
			Name: dto.Name,
			Age:  dto.Age,
		}
		users = append(users, user)
	}

	return &users, error
}
