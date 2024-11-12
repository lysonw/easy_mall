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
	Pid           uint64          // 商品id
	PCode         string          // 商品编号
	Name          string          // 商品名称
	Color         uint8           // 商品颜色
	Size          string          // 商品尺寸
	ImagPath      string          // 图片路径
	Title         string          // 商品标题
	SubTitle      string          // 商品子标题
	Description   string          // 商品描述
	Price         decimal.Decimal // 商品价格
	SaleTime      string          // 售卖时间
	ProductStatus uint8           // 商品生产状态 1:现货 2:非现货
	Status        uint8           // 状态 1:上架 2:下架 3:预售
	Inventory     int             // 库存数量
}
