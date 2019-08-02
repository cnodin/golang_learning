package user

import (
	. "demo7/handler"
	"demo7/pkg/errno"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	//infos, count, err := service.
}
