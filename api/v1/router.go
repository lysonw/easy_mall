package v1

import (
	"easy_mall/api/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		home := v1.Group("home")
		{
			home.GET("rotate", controller.RotateList)
			home.GET("new", controller.NewProductsOnline)
		}

	}
	return r
}
