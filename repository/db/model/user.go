package model

import "gorm.io/gorm"

type User struct {
	Uid      string // 唯一标识
	Nickname string // 昵称
	Phone    string // 电话
	Mail     string // 邮箱
	Password string // 密码
	gorm.Model
}

func (u *User) TableName() string {
	return "user"
}
