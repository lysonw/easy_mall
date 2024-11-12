package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResData 接口统一返回格式
type ResData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorReturn(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, &ResData{
		Code:    code,
		Message: message,
	})
	c.Abort()
}

func SuccessReturn(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResData{
		Code:    200,
		Message: "success",
		Data:    data,
	})
	c.Abort()
}
