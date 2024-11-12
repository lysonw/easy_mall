package request

import "easy_mall/model"

type RotateListReq struct {
	Num int `form:"num" json:"num"`
}

type NewProductReq struct {
	PCode string `form:"p_code" json:"p_code"`
	model.PageParam
}
