package model

import "gorm.io/gorm"

type ClickRecord struct {
	ID        int64  `json:"id" gorm:"id"`
	Uid       string `json:"uid" gorm:"uid"`               // 用户唯一标识
	Pid       int64  `json:"pid" gorm:"pid"`               // 商品ID
	PCode     string `json:"p_code" gorm:"p_code"`         // 商品编号
	ClickTime int64  `json:"click_time" gorm:"click_time"` // 点击时间
	gorm.Model
}

func (c *ClickRecord) TableName() string {
	return "click_record"
}
