package infra

import (
	"errors"
	"time"
)

type UserDto struct {
	Name string
	Age int
	CreatedAt time.Time
}

func GetUserDtos() ([] UserDto, error ) {
	if time.Now().Second() % 10 == 0 {
		return nil, errors.New("DB Error")
	}

	userDtos := make([]UserDto, 3)
	for i := range userDtos {
		userDtos[i] = UserDto{
			Name: "hoge",
			Age: i+ 20,
			CreatedAt:time.Now(),
		}
	}
	return userDtos, nil
}