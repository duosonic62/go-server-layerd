package controler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type ApiError interface {
	Error() string
	GetCode() int
	GetMessage() string
}

type ApplicationError struct {
	Message string
	Code    int
}

func New(code int, message string) ApplicationError {
	return ApplicationError{
		Code:    code,
		Message: message,
	}
}

func (appError ApplicationError) Error() string {
	return appError.Message
}

func (appError ApplicationError) GetCode() int {
	return appError.Code
}

func (appError ApplicationError) GetMessage() string {
	errors.New()
	return appError.Message
}

func commonError(context *gin.Context, err error) {
	switch e := err.(type) {
	case *ApplicationError:
		context.JSON(e.Code, err)
	default:
		context.JSON(500, New(500, e.Error()))
	}
}
