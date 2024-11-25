package dao

import (
	"context"
	i "easy_mall/init"
	base "easy_mall/model"
	"easy_mall/repository/db/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{i.DBClient(ctx)}
}

func (dao *ProductDao) Detail(pid string) (detail model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("pid = ?", pid).First(&detail).Error
	if err != nil {
		return
	}
	return
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

func (dao *ProductDao) CategoryList(start, end string) (list []model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Select("p_code,imag_path,title,sub_title,sale_time").
		Where("sale_time >= ? AND sale_time <= ?", start, end).
		Group("p_code,imag_path,title,sub_title,sale_time").
		Order("sale_time DESC").
		Find(&list).Error
	if err != nil {
		return
	}

	return
}
