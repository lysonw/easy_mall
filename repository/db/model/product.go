package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

/*
CREATE TABLE `products` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', -- 主键ID
  `pid` BIGINT(20) UNSIGNED NOT NULL COMMENT '商品id', -- 商品id
  `p_code` VARCHAR(255) NOT NULL COMMENT '商品编号', -- 商品编号
  `name` VARCHAR(255) NOT NULL COMMENT '商品名称', -- 商品名称
  `imag_path` VARCHAR(255) NOT NULL COMMENT '图片路径', -- 图片路径
  `title` VARCHAR(255) NOT NULL COMMENT '商品标题', -- 商品标题
  `sub_title` VARCHAR(255) NOT NULL COMMENT '商品子标题', -- 商品子标题
  `description` TEXT NOT NULL COMMENT '商品描述', -- 商品描述
  `price` DECIMAL(10, 2) NOT NULL COMMENT '商品价格', -- 商品价格
  `sale_time` DATETIME NOT NULL COMMENT '售卖时间', -- 售卖时间
  `status` TINYINT(3) UNSIGNED NOT NULL COMMENT '商品状态：1上架，2下架，3预售', -- 商品状态
  `inventory` INT NOT NULL COMMENT '库存', -- 库存
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间', -- 创建时间
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', -- 更新时间
  `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间', -- 删除时间
  PRIMARY KEY (`id`) -- 主键
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品信息表'; -- 表注释
*/

type Product struct {
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
	gorm.Model
}
