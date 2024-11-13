package controller

import (
	"easy_mall/model/request"
	"easy_mall/pkg/code"
	"easy_mall/pkg/util"
	"easy_mall/service"
	"github.com/gin-gonic/gin"
	"log"
)

func ProductDetail(c *gin.Context) {
	var param request.ProductDetailReq
	err := c.ShouldBind(&param)
	if err != nil {
		log.Println(err)
		util.ErrorReturn(c, code.ParamError, code.ParamErrorMsg)
		return
	}

	res, err := new(service.ProductService).GetProductDetail(c, param)
	if err != nil {
		log.Println(err)
		util.ErrorReturn(c, code.ServiceProductDetailError, code.ServiceProductDetailErrorMsg)
		return
	}

	util.SuccessReturn(c, res)
}

func ProductList(c *gin.Context) {

}

func ProductCategoryList(c *gin.Context) {

}
