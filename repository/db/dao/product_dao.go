package dao

import (
	"context"
	"easy_mall/init"
	base "easy_mall/model"
	"easy_mall/repository/db/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{init.DBClient(ctx)}
}

func (dao *ProductDao) List(condition map[string]interface{}, param base.PageParam) (list []model.Product, total int64, err error) {
	query := dao.DB.Model(&model.Product{}).Where(condition)

	if err = query.Count(&total).Error; err != nil || total == 0 {
		return
	}

	offset := (param.Page - 1) * param.PageSize

	err = query.Limit(param.PageSize).Offset(offset).Find(&list).Error
	if err != nil {
		return
	}

	return
}
