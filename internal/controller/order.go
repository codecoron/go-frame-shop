package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-frame-shop/api/backend"
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

func (c *cOrder) List(ctx context.Context, req *backend.OrderListReq) (res *backend.OrderListRes, err error) {
	orderListInput := model.OrderListInput{}
	if err = gconv.Struct(req, &orderListInput); err != nil {
		return nil, err
	}

	orderListOutput, err := service.Order().List(ctx, orderListInput)
	if err != nil {
		return nil, err
	}

	return &backend.OrderListRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  orderListOutput.List,
			Total: orderListOutput.Total,
			Page:  orderListOutput.Page,
			Size:  orderListOutput.Size,
		},
	}, err
}

func (c *cOrder) Detail(ctx context.Context, req *backend.OrderDetailReq) (res *backend.OrderDetailRes, err error) {
	in := model.OrderDetailInput{Id: req.Id}
	out, err := service.Order().Detail(ctx, in)
	if err != nil {
		return nil, err
	}
	res = &backend.OrderDetailRes{}
	err = gconv.Struct(out, res)
	return res, err
}
