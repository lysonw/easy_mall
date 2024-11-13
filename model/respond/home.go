package respond

import (
	"easy_mall/model"
	"github.com/shopspring/decimal"
)

type RotateListResp struct {
	Pid      uint64 `json:"pid"`
	ImagPath string `json:"imag_path"`
}

type NewProductResp struct {
	model.PageResp
	List []ProductInfo `json:"list"`
}
type ProductInfo struct {
	Pid           uint64          `json:"pid"`            // 商品id
	PCode         string          `json:"PCode"`          // 商品编号
	Name          string          `json:"name"`           // 商品名称
	Color         uint8           `json:"color"`          // 商品颜色
	Size          string          `json:"size"`           // 商品尺寸
	ImagPath      string          `json:"imag_path"`      // 图片路径
	Title         string          `json:"title"`          // 商品标题
	SubTitle      string          `json:"sub_title"`      // 商品子标题
	Description   string          `json:"description"`    // 商品描述
	Price         decimal.Decimal `json:"price"`          // 商品价格
	SaleTime      string          `json:"sale_time"`      // 售卖时间
	ProductStatus uint8           `json:"product_status"` // 商品生产状态 1:现货 2:非现货
	Status        uint8           `json:"status"`         // 状态 1:上架 2:下架 3:预售
	Inventory     int             `json:"inventory"`      // 库存数量
}

type HotProductResp struct {
	model.PageResp
	List []ProductInfo `json:"list"`
}
