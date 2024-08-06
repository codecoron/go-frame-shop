package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-frame-shop/api/frontend"
	"go-frame-shop/internal/model"
	"go-frame-shop/internal/service"
)

type cOrder struct {
}

var Order = cOrder{}

func (c *cOrder) Add(ctx context.Context, req *frontend.AddOrderReq) (res *frontend.AddOrderRes, err error) {

	in := model.OrderAddInput{}
	err = gconv.Scan(req, &in)
	if err != nil {
		return nil, err
	}

	addRes, err := service.Order().Add(ctx, in)
	if err != nil {
		return nil, err
	}

	return &frontend.AddOrderRes{
		Id: addRes.Id,
	}, err
}
