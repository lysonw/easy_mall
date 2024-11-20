package request

import base "easy_mall/model"

type ProductDetailReq struct {
	Pid uint64 `form:"pid" json:"pid" binding:"required"` // 商品id
}

type ProductCategoryListReq struct {
	Time string `json:"time"` // 时间 YYYY-MM-DD
}

type ProductListReq struct {
	PCode string `form:"pid" json:"p_code" binding:"required"`
	base.PageParam
}
