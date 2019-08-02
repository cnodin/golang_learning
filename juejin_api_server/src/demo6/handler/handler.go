package handler

import (
	"demo6/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

//resonse struct
type Response struct {
	Code int					`json:"code"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{})  {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code: code,
		Message: message,
		Data: data,
	})
}
