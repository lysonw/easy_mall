package model

import "gorm.io/gorm"

type ClickRecord struct {
	Uid       string // 唯一标识
	Pid       uint64 // 商品id
	PCode     string // 商品编号
	ClickTime string // 点击时间
	gorm.Model
}
