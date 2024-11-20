package respond

import (
	"easy_mall/model"
	"github.com/shopspring/decimal"
)

type RotateListResp struct {
	Pid      string `json:"pid"`
	ImagPath string `json:"imag_path"`
}

type NewProductResp struct {
	model.PageResp
	List []ProductInfo `json:"list"`
}
type ProductInfo struct {
	ID             int64           `json:"id"`
	Pid            string          `json:"pid"`             // 商品id
	PCode          string          `json:"PCode"`           // 商品编号
	Name           string          `json:"name"`            // 商品名称
	Color          uint8           `json:"color"`           // 商品颜色
	Size           string          `json:"size"`            // 商品尺寸
	ImagPath       string          `json:"imag_path"`       // 图片路径
	Title          string          `json:"title"`           // 商品标题
	SubTitle       string          `json:"sub_title"`       // 商品子标题
	Description    string          `json:"description"`     // 商品描述
	Price          decimal.Decimal `json:"price"`           // 商品价格
	SaleTime       string          `json:"sale_time"`       // 售卖时间
	AlbumPics      string          `json:"album_pics"`      // 相册图片
	PromotionPrice decimal.Decimal `json:"promotion_price"` // 促销价格
	OriginalPrice  decimal.Decimal `json:"original_price"`  // 市场价
	Inventory      int64           `json:"inventory"`       // 库存
	Sale           int64           `json:"sale"`            // 销量
	StockStatus    uint8           `json:"stock_status"`    // 是否现货
}

type HotProductResp struct {
	model.PageResp
	List []ProductInfo `json:"list"`
}
