package respond

import (
	base "easy_mall/model"
)

type ProductDetailResp struct {
	ProductInfo
}

type ProductListResp struct {
	base.PageResp
	List []ProductInfo `json:"list"`
}

type ProductCategoryListResp struct {
	List []CategoryList `json:"list"`
}

type CategoryList struct {
	PCode    string `json:"p_code"`    // 商品编号
	ImagPath string `json:"imag_path"` // 图片路径
	Title    string `json:"title"`     // 商品标题
	SubTitle string `json:"sub_title"` // 商品子标题
}
