package model

import "easy_mall/consts"

type PageParam struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (p PageParam) Verify() PageParam {
	if p.Page == 0 {
		p.Page = consts.PageDefault
	}

	if p.PageSize == 0 {
		p.PageSize = consts.PageSizeDefault
	} else if p.PageSize > consts.PageSizeMax {
		p.PageSize = consts.PageSizeMax
	}

	return p
}

type PageResp struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

type ProductClickRecord struct {
	Uid       string `json:"uid"`
	Pid       string `json:"pid"`
	PCode     string `json:"p_code"`
	ClickTime int64  `json:"click_time"`
}
