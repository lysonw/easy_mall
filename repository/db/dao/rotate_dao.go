package dao

import (
	"context"
	"easy_mall/init"
	"easy_mall/repository/db/model"
	"gorm.io/gorm"
	"log"
)

type RotateDao struct {
	*gorm.DB
}

func NewRotateDao(ctx context.Context) *RotateDao {
	return &RotateDao{init.DBClient(ctx)}
}

func (dao *RotateDao) List(limit int) (list []model.Rotate, err error) {
	err = dao.DB.Model(&model.Rotate{}).Select("imag_path,pid").Limit(limit).Find(&list).Error
	if err != nil {
		log.Printf("get rotate list err:%s", err.Error())
	}
	return
}
