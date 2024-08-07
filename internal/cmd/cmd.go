package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"go-frame-shop/internal/controller"
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
					)
				})
			})

			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
				)
				group.Bind(
					controller.User.Register, //用户注册
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Bind(
						controller.User.UpdatePassword, //当前用户修改密码
						controller.Order.Add,
						controller.OrderGoodsComments.Add,
						controller.OrderGoodsComments.Delete,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
