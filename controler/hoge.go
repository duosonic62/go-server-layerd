package controler

import "errors"

type User struct {
	Name string
	Age  int
}



func NewUser(name string, age int) (*User, error) {
	if len(name) == 0 {
		return nil, errors.New("bad request")
	}

	user := User{
		Name: name,
		Age:  age,
	}

	return &user, nil
}
