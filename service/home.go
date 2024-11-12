package service

import (
	"context"
	"easy_mall/model/request"
	"easy_mall/model/respond"
	"easy_mall/repository/db/dao"
	"log"
	"time"
)

type Home struct {
}

func (h *Home) RotateList(ctx context.Context, limit int) (res []respond.RotateListResp, err error) {
	list, err := dao.NewRotateDao(ctx).List(limit)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range list {
		var item respond.RotateListResp
		item.Pid = v.Pid
		item.ImagPath = v.ImagPath
		res = append(res, item)
	}

	return
}

func (h *Home) NewProductList(ctx context.Context, req request.NewProductReq) (res respond.NewProductResp, err error) {
	now := time.Now().Format(time.DateOnly)
	condition := make(map[string]interface{})
	condition["sale_time"] = now
	if req.PCode != "" {
		condition["p_code"] = req.PCode
	}

	list, total, err := dao.NewProductDao(ctx).List(condition, req.PageParam)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range list {
		item := respond.ProductInfo{
			Pid:           v.Pid,
			PCode:         v.PCode,
			Name:          v.Name,
			Color:         v.Color,
			Size:          v.Size,
			ImagPath:      v.ImagPath,
			Title:         v.Title,
			SubTitle:      v.SubTitle,
			Description:   v.Description,
			Price:         v.Price,
			SaleTime:      v.SaleTime,
			ProductStatus: v.ProductStatus,
			Status:        v.Status,
			Inventory:     v.Inventory,
		}
		res.List = append(res.List, item)
	}

	res.Total = total
	res.Page = req.PageParam.Page
	res.PageSize = req.PageParam.PageSize
	return
}
