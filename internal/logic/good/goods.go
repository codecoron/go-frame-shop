package good

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go-frame-shop/internal/dao"
	"go-frame-shop/internal/model"
	"go-frame-shop/internal/model/entity"
	"go-frame-shop/internal/service"
)

type sGoods struct{}

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

func (s *sGoods) Create(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error) {
	lastInsertID, err := dao.GoodsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.GoodsCreateOutput{Id: uint(lastInsertID)}, err
}

func (s *sGoods) GetList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error) {
	var (
		m = dao.GoodsInfo.Ctx(ctx)
	)

	out = &model.GoodsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)
	// todo 又是scan了两遍
	var list []*entity.GoodsInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}

	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// Delete 删除
func (s *sGoods) Delete(ctx context.Context, id uint) (err error) {
	_, err = dao.GoodsInfo.Ctx(ctx).Where(
		g.Map{dao.GoodsInfo.Columns().Id: id}).Delete()
	if err != nil {
		return err
	}
	return
}

func (s *sGoods) Update(ctx context.Context, in model.GoodsUpdateInput) error {
	_, err := dao.GoodsInfo.Ctx(ctx).Data(in).FieldsEx(dao.
		GoodsInfo.Columns().Id).Where(dao.GoodsInfo.Columns().Id, in.Id).Update()
	return err
}
