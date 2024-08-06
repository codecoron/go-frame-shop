package model

// AddOrderGoodsCommentsInput 没有记录评价userid,可以根据订单ID找到userid。根据userid找评论，不是高频场景
type AddOrderGoodsCommentsInput struct {
	OrderId        uint
	GoodsId        uint
	GoodsOptionsId uint
	ParentId       uint
	Content        string
}

type AddOrderGoodsCommentsOutput struct {
	Id uint
}

type DelOrderGoodsCommentsInput struct {
	Id uint
}

type DelOrderGoodsCommentsOutput struct {
	Id uint
}
