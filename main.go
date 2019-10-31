package main

import (
	"fmt"
	"sample-sever/controler"
)

func main() {
	user, error := controler.NewUser("sample", 10)
	if error != nil {
		fmt.Println("error!")
		fmt.Println(error)
	}
	fmt.Println(*user)
}