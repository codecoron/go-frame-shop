package backend

//type OrderListReq struct {
//	g.Meta `path:"/order/list" tags:"订单列表" method:"get" summary:"订单列表"`
//	backend.CommonPaginationReq
//	Number         string `json:"number"           dc:"订单编号"`
//	UserId         int    `json:"userId"           dc:"用户id"`
//	PayType        int    `json:"payType"          dc:"支付方式 1微信 2支付宝 3云闪付"`
//	PayAtGte       string `json:"payAtGte"         dc:"支付时间>="`
//	PayAtLte       string `json:"payAtLte"         dc:"支付时间<="`
//	Status         int    `json:"status"           dc:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价"`
//	ConsigneePhone string `json:"consigneePhone"   dc:"收货人手机号"`
//	PriceGte       int    `json:"priceGte"         dc:"订单金额>= 单位分"`
//	PriceLte       int    `json:"priceLte"         dc:"订单金额<= 单位分"`
//	DateGte        string `json:"dateGte"          dc:"創建时间>="`
//	DateLte        string `json:"dateLte"          dc:"創建时间<="`
//}
//
//type OrderListRes struct {
//	backend.CommonPaginationRes
//}
