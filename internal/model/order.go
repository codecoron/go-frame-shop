package model

type OrderAddInput struct {
	UserId           uint
	Number           string
	Remark           string `description:"备注"`
	Price            int    `description:"订单金额 单位分"`
	CouponPrice      int    `description:"优惠券金额 单位分"`
	ActualPrice      int    `description:"实际支付金额 单位分"`
	ConsigneeName    string `description:"收货人姓名"`
	ConsigneePhone   string `description:"收货人手机号"`
	ConsigneeAddress string `description:"收货人详细地址"`
	//商品订单维度
	OrderAddGoodsInfos []*OrderAddGoodsInfo
}

type OrderAddGoodsInfo struct {
	Id             int
	OrderId        int
	GoodsId        int
	GoodsOptionsId int
	Count          int
	Remark         string
	Price          int
	CouponPrice    int
	ActualPrice    int
}

type OrderAddOutput struct {
	Id uint `json:"id"`
}
