package controler

import (
	"github.com/gin-gonic/gin"
	"sample-sever/domain/errors"
)

// アプリに表示するエラー
type PresentationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// API共通のエラーハンドリング
func commonError(context *gin.Context, err error) {
	switch e := err.(type) {
	case errors.ApplicationError:
		context.JSON(e.GetCode(), PresentationError{Code: e.GetCode(), Message: e.GetFrontMessage()})
	default:
		context.JSON(500, PresentationError{Code: 500, Message: "INTERNAL SERVER ERROR"})
	}
}
