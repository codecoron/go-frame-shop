package model

import (
	"go-frame-shop/internal/model/do"
	"go-frame-shop/internal/model/entity"
)

type RegisterInput struct {
	Name         string `json:"name"         description:"用户名" v:"required#用户名必填"`
	Password     string `json:"password"     description:"密码" v:"password"`
	Avatar       string `json:"avatar"       description:"头像"`
	UserSalt     string `json:"userSalt"     description:"加密盐 生成密码用"`
	Sex          int    `json:"sex"          description:"1男 2女"`
	Status       int    `json:"status"       description:"1正常 2拉黑冻结"`
	Sign         string `json:"sign"         description:"个性签名"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type RegisterOutput struct {
	Id uint `json:"id"`
}

// UpdatePasswordInput 修改密码
type UpdatePasswordInput struct {
	Password     string `json:"password"  v:"password"   description:""`
	UserSalt     string `json:"userSalt,omitempty"     description:"加密盐 生成密码用"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type UpdatePasswordOutput struct {
	Id uint `json:"id"`
}

type OrderListInput struct {
	Page           int    // 分页号码
	Size           int    // 分页数量，最大50
	Number         string `json:"number"           dc:"订单编号"`
	UserId         int    `json:"userId"           dc:"用户id"`
	PayType        int    `json:"payType"          dc:"支付方式 1微信 2支付宝 3云闪付"`
	PayAtGte       string `json:"payAtGte"         dc:"支付时间>="`
	PayAtLte       string `json:"payAtLte"         dc:"支付时间<="`
	Status         int    `json:"status"           dc:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价"`
	ConsigneePhone string `json:"consigneePhone"   dc:"收货人手机号"`
	PriceGte       int    `json:"priceGte"         dc:"订单金额>= 单位分"`
	PriceLte       int    `json:"priceLte"         dc:"订单金额<= 单位分"`
	DateGte        string `json:"dateGte"          dc:"創建时间>="`
	DateLte        string `json:"dateLte"          dc:"創建时间<="`
}

type OrderListOutput struct {
	List  []OrderListOutputItem
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type OrderListOutputItem struct {
	entity.OrderInfo
}

type OrderDetailInput struct {
	Id uint
}

type OrderDetailOutput struct {
	//entity.OrderInfo
	do.OrderInfo
	GoodsInfo []*do.OrderGoodsInfo `orm:"with:order_id=id"`
}
