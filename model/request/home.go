package request

import "easy_mall/model"

type RotateListReq struct {
	Num int `form:"num" json:"num"`
}

type NewProductReq struct {
	model.PageParam
}

type HotProductReq struct {
	model.PageParam
}
