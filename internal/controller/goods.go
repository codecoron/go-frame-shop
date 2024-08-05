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
