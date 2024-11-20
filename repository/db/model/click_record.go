package model

import "gorm.io/gorm"

/*
CREATE TABLE `click_record` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `pid` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `p_code` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `click_time` bigint NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_click_record_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/

type ClickRecord struct {
	ID        int64  `json:"id" gorm:"id"`
	Uid       string `json:"uid" gorm:"uid"`               // 用户唯一标识
	Pid       string `json:"pid" gorm:"pid"`               // 商品ID
	PCode     string `json:"p_code" gorm:"p_code"`         // 商品编号
	ClickTime int64  `json:"click_time" gorm:"click_time"` // 点击时间
	gorm.Model
}

func (c *ClickRecord) TableName() string {
	return "click_record"
}
