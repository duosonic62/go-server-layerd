package controler

import (
	"github.com/gin-gonic/gin"
	"sample-sever/application"
)

func GetUsers(context *gin.Context) {
	users, error := application.GetUsers()
	if error != nil {
		commonError(context, error)
	}else {
		context.JSON(200, users)
	}
}
