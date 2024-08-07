// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CouponInfo is the golang structure of table coupon_info for DAO operations like Where/Data.
type CouponInfo struct {
	g.Meta     `orm:"table:coupon_info, do:true"`
	Id         interface{} //
	Name       interface{} //
	Price      interface{} // 优惠前面值 单位分
	GoodsId    interface{} // goods_id
	CategoryId interface{} // 关联使用的分类id
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
}
