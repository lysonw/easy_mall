package dao

import (
	"context"
	i "easy_mall/init"
	"easy_mall/repository/db/model"
	"gorm.io/gorm"
	"log"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{i.DBClient(ctx)}
}

func (dao *UserDao) Insert(info model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(info).Error
	if err != nil {
		log.Printf("insert user info err:%s", err.Error())
	}
	return
}

func (dao *UserDao) Get(uid string) (info model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("uid = ?", uid).First(&info).Error
	if err != nil {
		log.Printf("get user info err:%s", err.Error())
	}
	return
}
