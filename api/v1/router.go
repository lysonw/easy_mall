package v1

import (
	"easy_mall/api/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		// 暂时先不做登录这一块的内容
		user := v1.Group("user")
		{
			user.GET("login")     // 登录
			user.POST("register") // 注册
		}

		home := v1.Group("home")
		{
			home.GET("rotate", controller.RotateList)        // 首页轮播
			home.GET("new", controller.NewProductsOnline)    // 新品商品推荐
			home.GET("hot", controller.HotProductsRecommend) // 热点商品推荐
		}

		product := v1.Group("product")
		{
			product.GET("detail", controller.ProductDetail)         // 商品详情
			product.GET("list", controller.ProductList)             // 商品列表
			product.GET("category", controller.ProductCategoryList) // 分类列表
		}

	}
	return r
}
