package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-frame-shop/api/frontend"
	"go-frame-shop/internal/model"
	"go-frame-shop/internal/service"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *frontend.RegisterReq) (res *frontend.RegisterRes, err error) {

	in := model.RegisterInput{}
	err = gconv.Scan(req, &in)
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, in)
	if err != nil {
		return nil, err
	}

	return &frontend.RegisterRes{Id: out.Id}, err
}

func (c *cUser) UpdatePassword(ctx context.Context, req *frontend.UpdatePasswordReq) (res *frontend.UpdatePasswordRes, err error) {
	data := model.UpdatePasswordInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().UpdatePassword(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.UpdatePasswordRes{Id: out.Id}, nil
}
