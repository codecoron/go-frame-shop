package frontend

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-frame-shop/api/frontend"
	"go-frame-shop/internal/consts"
	"go-frame-shop/internal/model"
	"go-frame-shop/internal/service"
)

// Article 内容管理
var Article = cArticle{}

type cArticle struct{}

func (a *cArticle) Create(ctx context.Context, req *frontend.ArticleAddReq) (res *frontend.ArticleAddRes, err error) {
	data := model.ArticleCreateInput{}
	//这里不需要用scan 用struct就可以 因为不涉及到嵌套，就是最简单的结构体转换
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	data.IsAdmin = consts.ArticleIsUser
	data.UserId = gconv.Int(ctx.Value(consts.CtxUserId))
	out, err := service.Article().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.ArticleAddRes{Id: out.Id}, nil
}
