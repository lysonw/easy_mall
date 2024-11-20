package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

/*
CREATE TABLE `product` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `pid` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `p_code` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `color` tinyint unsigned NOT NULL,
  `size` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `imag_path` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `album_pics` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `title` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sub_title` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `promotion_price` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `original_price` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sale_time` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `inventory` bigint NOT NULL,
  `sort` tinyint unsigned NOT NULL,
  `sale` bigint NOT NULL,
  `stock_status` tinyint unsigned NOT NULL,
  `delete_status` tinyint unsigned NOT NULL,
  `publish_status` tinyint unsigned NOT NULL,
  `new_status` tinyint unsigned NOT NULL,
  `recommend_status` tinyint unsigned NOT NULL,
  `rotate_status` tinyint unsigned NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_product_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/

type Product struct {
	ID              int64           `json:"id" gorm:"id"`
	Pid             string          `json:"pid" gorm:"pid"`                           // 商品id
	PCode           string          `json:"p_code" gorm:"p_code"`                     // 商品编号
	Name            string          `json:"name" gorm:"name"`                         // 商品名称
	Color           uint8           `json:"color" gorm:"color"`                       // 商品颜色
	Size            string          `json:"size" gorm:"size"`                         // 商品尺寸
	ImagPath        string          `json:"imag_path" gorm:"imag_path"`               // 图片路径
	AlbumPics       string          `json:"album_pics" gorm:"album_pics"`             // 画册图片，连产品图片限制为5张，以逗号分割
	Title           string          `json:"title" gorm:"title"`                       // 商品标题
	SubTitle        string          `json:"sub_title" gorm:"sub_title"`               // 商品子标题
	Description     string          `json:"description" gorm:"description"`           // 商品描述
	Price           decimal.Decimal `json:"price" gorm:"price"`                       // 价格
	PromotionPrice  decimal.Decimal `json:"promotion_price" gorm:"promotion_price"`   // 促销价格
	OriginalPrice   decimal.Decimal `json:"original_price" gorm:"price"`              // 市场价
	SaleTime        string          `json:"sale_time" gorm:"sale_time"`               // 售卖时间
	Inventory       int64           `json:"inventory" gorm:"inventory"`               // 库存数量
	Sort            uint8           `json:"sort" gorm:"sort"`                         // 排序
	Sale            int64           `json:"sale" gorm:"sale"`                         // 销量
	StockStatus     uint8           `json:"stock_status" gorm:"stock_status"`         // 是否是现货 0->不是；1->是
	DeleteStatus    uint8           `json:"delete_status" gorm:"delete_status"`       // 删除状态：0->未删除；1->已删除
	PublishStatus   uint8           `json:"publish_status" gorm:"publish_status"`     // 上架状态：0->下架；1->上架
	NewStatus       uint8           `json:"new_status" gorm:"new_status"`             // 新品状态:0->不是新品；1->新品
	RecommendStatus uint8           `json:"recommend_status" gorm:"recommend_status"` // 推荐状态；0->不推荐；1->推荐
	RotateStatus    uint8           `json:"rotate_status" gorm:"rotate_status"`       // 轮播状态；0->不轮播；1->轮播
	gorm.Model
}

func (p *Product) TableName() string {
	return "product"
}
