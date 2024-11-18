package respond

import base "easy_mall/model"

type ProductDetailResp struct {
	ProductInfo
}

type ProductListResp struct {
	base.PageResp
	List []ProductInfo `json:"list"`
}

type ProductCategoryListResp struct {
	List []ProductInfo `json:"list"`
}
