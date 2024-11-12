package v1

import (
	"easy_mall/api/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{

		v1.GET("login")     // 登录
		v1.POST("register") // 注册

		callback := v1.Group("callback")
		{
			callback.POST("click") // 点击事件回调
		}

		home := v1.Group("home")
		{
			home.GET("rotate", controller.RotateList)        // 首页轮播
			home.GET("new", controller.NewProductsOnline)    // 新品商品推荐
			home.GET("hot", controller.HotProductsRecommend) // 热点上坪推荐
		}

	}
	return r
}
