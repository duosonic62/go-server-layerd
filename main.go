package main

import (
	"github.com/gin-gonic/gin"
	"sample-sever/controler"
)

func main() {
	router := gin.Default()

	router.GET("/users", controler.GetUsers)

	router.Run()
}