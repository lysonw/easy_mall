package service

import (
	"context"
	base "easy_mall/model"
	"easy_mall/model/request"
	"easy_mall/model/respond"
	"easy_mall/repository/cache"
	"easy_mall/repository/db/dao"
	"encoding/json"
	"log"
	"time"
)

type ProductService struct{}

func (p *ProductService) GetProductDetail(ctx context.Context, req request.ProductDetailReq) (res respond.ProductDetailResp, err error) {

	detail, err := dao.NewProductDao(ctx).Detail(req.Pid)
	if err != nil {
		return
	}
	info := respond.ProductInfo{
		ID:             detail.ID,
		Pid:            detail.Pid,
		PCode:          detail.PCode,
		Name:           detail.Name,
		Color:          detail.Color,
		Size:           detail.Size,
		ImagPath:       detail.ImagPath,
		Title:          detail.Title,
		SubTitle:       detail.SubTitle,
		Description:    detail.Description,
		Price:          detail.Price,
		SaleTime:       detail.SaleTime,
		AlbumPics:      detail.AlbumPics,
		PromotionPrice: detail.PromotionPrice,
		OriginalPrice:  detail.OriginalPrice,
		Inventory:      detail.Inventory,
		Sale:           detail.Sale,
		StockStatus:    detail.StockStatus,
	}
	res.ProductInfo = info

	// 商品访问记录
	record := base.ProductClickRecord{
		Uid:       "",
		Pid:       detail.Pid,
		PCode:     detail.PCode,
		ClickTime: time.Now().Unix(),
	}
	b, err := json.Marshal(record)
	if err != nil {
		log.Println(err)
	}
	c := cache.Cache{Ctx: ctx}
	if err = c.LPush(cache.ProductClickRecordCacheKey, string(b)); err != nil {
		return
	}

	return
}
func (p *ProductService) GetProductList(ctx context.Context, req request.ProductListReq) (res respond.ProductListResp, err error) {
	condition := make(map[string]interface{})
	condition["p_code"] = req.PCode
	data, total, err := dao.NewProductDao(ctx).List(condition, req.PageParam)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range data {
		item := respond.ProductInfo{
			ID:             v.ID,
			Pid:            v.Pid,
			PCode:          v.PCode,
			Name:           v.Name,
			Color:          v.Color,
			Size:           v.Size,
			ImagPath:       v.ImagPath,
			Title:          v.Title,
			SubTitle:       v.SubTitle,
			Description:    v.Description,
			Price:          v.Price,
			SaleTime:       v.SaleTime,
			AlbumPics:      v.AlbumPics,
			PromotionPrice: v.PromotionPrice,
			OriginalPrice:  v.OriginalPrice,
			Inventory:      v.Inventory,
			Sale:           v.Sale,
			StockStatus:    v.StockStatus,
		}
		res.List = append(res.List, item)
	}

	res.PageResp.Total = total
	res.PageResp.Page = req.PageParam.Page
	res.PageResp.PageSize = req.PageParam.PageSize
	return
}

/*
	分类列表返回的结构 ：
			新品 （当月）
		    全部
			xxx年12月
		    xxx年11月
			xxx年10月
		    ... ...
*/

func (p *ProductService) GetProductCategoryList(ctx context.Context, req request.ProductCategoryListReq) (res respond.ProductCategoryListResp, err error) {
	var start, end string
	if start, end, err = getProductCategoryTimeRange(req.Time); err != nil {
		log.Println(err)
		return
	}

	data, err := dao.NewProductDao(ctx).CategoryList(start, end)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range data {
		item := respond.CategoryList{
			PCode:    v.PCode,
			ImagPath: v.ImagPath,
			Title:    v.Title,
			SubTitle: v.SubTitle,
		}
		res.List = append(res.List, item)
	}

	return
}
