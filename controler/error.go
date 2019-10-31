package controler

import "github.com/gin-gonic/gin"

type apiError struct {
	Message string
}

func commonError(context *gin.Context, err error) {
	context.JSON(500, err)
}