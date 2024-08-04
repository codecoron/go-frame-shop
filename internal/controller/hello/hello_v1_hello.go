package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "go-frame-shop/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}

func (c *ControllerV1) Hello2(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	// Implement your logic here
	g.RequestFromCtx(ctx).Response.Writeln("Hello from Hello2!")

	// Example of constructing HelloRes
	res = &v1.HelloRes{
		// Populate with appropriate fields based on your logic
	}

	return res, nil
}
