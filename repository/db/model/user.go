package model

import "gorm.io/gorm"

type User struct {
	Uid      string `json:"uid"`      // 唯一标识
	Nickname string `json:"nickname"` // 昵称
	Phone    string `json:"phone"`    // 电话
	Mail     string `json:"mail"`     // 邮箱
	Password string `json:"password"` // 密码
	gorm.Model
}
