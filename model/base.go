package model

import "easy_mall/consts"

type PageParam struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func (p *PageParam) Verify() *PageParam {
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
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}
