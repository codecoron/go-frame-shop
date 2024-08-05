package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-frame-shop/api/backend"
	"go-frame-shop/internal/model"
	"go-frame-shop/internal/service"
)

var Goods = cGoods{}

type cGoods struct{}

func (a *cGoods) Create(ctx context.Context, req *backend.GoodsReq) (res *backend.GoodsRes, err error) {
	//return nil, err
	data := model.GoodsCreateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Goods().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsRes{Id: out.Id}, err
}

func (a *cGoods) List(ctx context.Context, req *backend.GoodsGetListCommonReq) (res *backend.GoodsGetListCommonRes, err error) {
	out, err := service.Goods().GetList(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.GoodsGetListCommonRes{List: out.List, Page: out.Page, Size: out.Size, Total: out.Total}, err
}

func (a *cGoods) Delete(ctx context.Context, req *backend.GoodsDeleteReq) (res *backend.GoodsDeleteRes, err error) {

	err = service.Goods().Delete(ctx, req.Id)

	return &backend.GoodsDeleteRes{}, err
}

func (a *cGoods) Update(ctx context.Context, req *backend.GoodsUpdateReq) (res *backend.GoodsUpdateRes, err error) {
	data := model.GoodsUpdateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	err = service.Goods().Update(ctx, data)

	return &backend.GoodsUpdateRes{Id: req.Id}, err
}
