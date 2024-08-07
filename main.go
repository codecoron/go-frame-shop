package main

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "go-frame-shop/internal/packed"

	_ "go-frame-shop/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"go-frame-shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
