package order

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"go-frame-shop/internal/consts"
	"go-frame-shop/internal/dao"
	"go-frame-shop/internal/model"
	"go-frame-shop/internal/service"
	"go-frame-shop/utility"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

func (s *sOrder) Add(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	in.Number = utility.GetOrderNum()
	out = &model.OrderAddOutput{}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 生成主订单
		lastInsertId, err := dao.OrderInfo.Ctx(
			ctx).InsertAndGetId(in)
		if err != nil {
			return err
		}
		for _, info := range in.OrderAddGoodsInfos {
			info.OrderId = gconv.Int(lastInsertId)
			_, err := dao.OrderGoodsInfo.Ctx(ctx).Insert(info)
			if err != nil {
				return err
			}
		}
		// 更新商品和销量库存
		for _, info := range in.OrderAddGoodsInfos {
			// 商品销量增加
			_, err := dao.GoodsInfo.Ctx(ctx).WherePri(info.GoodsId).Increment(dao.GoodsInfo.Columns().Sale, info.Count)
			if err != nil {
				return err
			}
			// 商品减少库存
			_, err2 := dao.GoodsInfo.Ctx(ctx).WherePri(info.GoodsId).Decrement(dao.GoodsInfo.Columns().Stock, info.Count)
			if err2 != nil {
				return err2
			}
			// 商品规格减少库存
			_, err3 := dao.GoodsInfo.Ctx(ctx).WherePri(info.GoodsId).Decrement(dao.GoodsInfo.Columns().Stock, info.Count)
			if err3 != nil {
				return err3
			}
		}
		out.Id = uint(lastInsertId)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s *sOrder) List(ctx context.Context, in model.OrderListInput) (out *model.OrderListOutput, err error) {
	out = &model.OrderListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	whereCondition := s.orderListCondition(in)
	m := dao.OrderInfo.Ctx(ctx).Where(whereCondition)

	if err = m.Page(in.Page, in.Size).Scan(&out.List); err != nil {
		return nil, err
	}

	if out.Total, err = m.Count(); err != nil {
		return nil, err
	}
	return
}

func (s *sOrder) orderListCondition(in model.OrderListInput) *gmap.Map {
	m := gmap.New()

	if in.Number != "" {
		m.Set(dao.OrderInfo.Columns().Number+" like ", "%"+in.Number+"%")
	}

	if in.UserId != 0 {
		m.Set(dao.OrderInfo.Columns().UserId, in.UserId)
	}

	if in.PayType != 0 {
		m.Set(dao.OrderInfo.Columns().PayType, in.PayType)
	}

	if in.PayAtGte != "" {
		m.Set(dao.OrderInfo.Columns().PayAt+" >= ", gtime.New(in.PayAtGte).StartOfDay())
	}

	if in.PayAtLte != "" {
		m.Set(dao.OrderInfo.Columns().PayAt+" <= ", gtime.New(in.PayAtLte).EndOfDay())
	}

	if in.Status != 0 {
		m.Set(dao.OrderInfo.Columns().Status, in.Status)
	}

	if in.ConsigneePhone != "" {
		m.Set(dao.OrderInfo.Columns().ConsigneePhone+" like ", "%"+in.ConsigneePhone+"%")
	}

	if in.PriceGte != 0 {
		m.Set(dao.OrderInfo.Columns().Price+" >= ", in.PriceGte)
	}

	if in.PriceLte != 0 {
		m.Set(dao.OrderInfo.Columns().Price+" <= ", in.PriceLte)
	}

	if in.DateGte != "" {
		m.Set(dao.OrderInfo.Columns().CreatedAt+" >= ", gtime.New(in.DateGte).StartOfDay())
	}

	if in.DateLte != "" {
		m.Set(dao.OrderInfo.Columns().CreatedAt+" <= ", gtime.New(in.DateLte).EndOfDay())
	}

	return m
}

func (s *sOrder) Detail(ctx context.Context, in model.OrderDetailInput) (out *model.OrderDetailOutput, err error) {
	err = dao.OrderInfo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	return
}
