package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"go-frame-shop/internal/controller"
	"go-frame-shop/internal/controller/frontend"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			//s.Group("/", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse)
			//	group.Bind(
			//		hello.NewV1(),
			//	)
			//})
			// 启动管理后台gtoken
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}

			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					//service.Middleware().CORS,
					//service.Middleware().Ctx,
					//service.Middleware().ResponseHandler,
					ghttp.MiddlewareHandlerResponse,
				)
				// 不需要登录的路由组绑定
				group.Bind(
					controller.Admin.Create, // 管理员
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.Admin.List,   // 管理员
						controller.Admin.Update, // 更新
						controller.Admin.Delete,
						controller.Goods, //商品管理
						controller.Order.List,
						controller.Order.Detail,
						controller.GoodsOptions, //商品规格管理
						controller.Category,     //商品分类
						controller.Coupon,       // 优惠卷
						controller.Role,         // 角色管理
						controller.Permission,   // 权限
						controller.Rotation,     // 轮播图
					)
				})
			})

			frontendToken, err := StartFrontendGToken()
			if err != nil {
				return err
			}

			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
				)
				group.Bind(
					controller.User.Register, //用户注册
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := frontendToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.User.UpdatePassword, //当前用户修改密码
						controller.Order.Add,
						controller.OrderGoodsComments.Add,
						controller.OrderGoodsComments.Delete,
						frontend.Article,
						controller.Cart,
						controller.Collection,
						controller.Comment, //评论相关

						frontend.Refund, // 退货相关
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
