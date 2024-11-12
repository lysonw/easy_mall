package controller

import (
	"easy_mall/model/request"
	"easy_mall/pkg/code"
	"easy_mall/pkg/util"
	"easy_mall/service"
	"github.com/gin-gonic/gin"
	"log"
)

func RotateList(c *gin.Context) {
	var param request.RotateListReq
	err := c.ShouldBind(&param)
	if err != nil {
		log.Println(err)
		util.ErrorReturn(c, code.ParamError, code.ParamErrorMsg)
		return
	}

	res, err := new(service.Home).RotateList(c, param.Num)
	if err != nil {
		log.Println(err)
		util.ErrorReturn(c, code.ServiceRotateError, code.ServiceRotateErrorMsg)
		return
	}

	util.SuccessReturn(c, res)
}

func NewProductsOnline(c *gin.Context) {
	var param request.NewProductReq
	err := c.ShouldBind(&param)
	if err != nil {
		log.Println(err)
		util.ErrorReturn(c, code.ParamError, code.ParamErrorMsg)
		return
	}

	res, err := new(service.Home).NewProductList(c, param)
	if err != nil {
		log.Println(err)
		util.ErrorReturn(c, code.ServiceRotateError, code.ServiceRotateErrorMsg)
		return
	}

	util.SuccessReturn(c, res)
}

func HotProductsRecommend(c *gin.Context) {

}
