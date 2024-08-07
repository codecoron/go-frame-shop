package article

import (
	"context"
	"go-frame-shop/internal/dao"
	"go-frame-shop/internal/model"
	"go-frame-shop/internal/service"
)

type sArticle struct{}

func init() {
	service.RegisterArticle(New())
}

func New() *sArticle {
	return &sArticle{}
}

func (s *sArticle) Create(ctx context.Context, in model.ArticleCreateInput) (out model.ArticleCreateOutput, err error) {
	lastInsertID, err := dao.ArticleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.ArticleCreateOutput{Id: uint(lastInsertID)}, err
}
