package dao

import (
	"context"
	i "easy_mall/init"
	"easy_mall/repository/db/model"
	"gorm.io/gorm"
)

type ClickRecordDao struct {
	*gorm.DB
}

func NewClickRecordDao(ctx context.Context) *ClickRecordDao {
	return &ClickRecordDao{i.DBClient(ctx)}
}

func (dao *ClickRecordDao) Insert(data model.ClickRecord) (err error) {
	err = dao.DB.Model(&model.ClickRecord{}).Create(&data).Error
	if err != nil {
		return
	}
	return
}

func (dao *ClickRecordDao) BatchInsert(data []model.ClickRecord, batchSize int) (err error) {
	err = dao.DB.Model(&model.ClickRecord{}).CreateInBatches(data, batchSize).Error
	if err != nil {
		return
	}
	return
}
