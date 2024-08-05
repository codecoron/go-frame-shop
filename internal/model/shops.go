package model

import "go-frame-shop/internal/model/entity"

// GoodsCreateUpdateBase 创建/修改内容基类
type GoodsCreateUpdateBase struct {
	PicUrl           string
	Name             string
	Price            int
	Level1CategoryId int
	Level2CategoryId int
	Level3CategoryId int
	Brand            string
	Stock            int
	Sale             int
	Tags             string
	DetailInfo       string
}

// GoodsCreateInput 创建内容
type GoodsCreateInput struct {
	GoodsCreateUpdateBase
}

// GoodsCreateOutput 创建内容返回结果
type GoodsCreateOutput struct {
	Id uint `json:"id"`
}

type GoodsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

type GoodsGetListOutput struct {
	List  []GoodsGetListOutputItem
	Page  int
	Size  int
	Total int
}

// GoodsGetListOutputItem logic的model引用了entity
type GoodsGetListOutputItem struct {
	entity.GoodsInfo
}
