package service

import (
	"context"
	base "easy_mall/model"
	"easy_mall/model/request"
	"easy_mall/model/respond"
	"easy_mall/repository/cache"
	"easy_mall/repository/db/dao"
	"time"
)

type ProductService struct{}

func (p *ProductService) GetProductDetail(ctx context.Context, req request.ProductDetailReq) (res respond.ProductDetailResp, err error) {

	detail, err := dao.NewProductDao(ctx).Detail(req.Pid)
	if err != nil {
		return
	}
	info := respond.ProductInfo{
		Pid:           detail.Pid,
		PCode:         detail.PCode,
		Name:          detail.Name,
		Color:         detail.Color,
		Size:          detail.Size,
		ImagPath:      detail.ImagPath,
		Title:         detail.Title,
		SubTitle:      detail.SubTitle,
		Description:   detail.Description,
		Price:         detail.Price,
		SaleTime:      detail.SaleTime,
		ProductStatus: detail.ProductStatus,
		Status:        detail.Status,
		Inventory:     detail.Inventory,
	}
	res.ProductInfo = info

	// 商品访问记录
	record := base.ProductClickRecord{
		Uid:       "",
		Pid:       detail.Pid,
		PCode:     detail.PCode,
		ClickTime: time.Now().Format(time.DateTime),
	}
	c := cache.Cache{Ctx: ctx}
	if err = c.LPush(cache.ProductClickRecordCacheKey, record); err != nil {
		return
	}

	return
}
