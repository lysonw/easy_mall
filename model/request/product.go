package request

import base "easy_mall/model"

type ProductDetailReq struct {
	Pid string `json:"pid" form:"pid" binding:"required"` // 商品id
}

type ProductCategoryListReq struct {
	Time string `json:"time" form:"time"` // 时间 YYYY-MM-DD
}

type ProductListReq struct {
	PCode string `json:"p_code" form:"p_code" binding:"required"`
	base.PageParam
}
