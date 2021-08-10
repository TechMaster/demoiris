package controller

import "github.com/kataras/iris/v12"

func Handler1(ctx iris.Context) {
	ctx.ViewData("key1", "Value from func Handler1(ctx iris.Context)")
	ctx.Next()
}

func Handler2(ctx iris.Context) {
	ctx.ViewData("key2", "Value from func Handler2(ctx iris.Context)")
	ctx.Next()
}

func Handler3(ctx iris.Context) {
	ctx.ViewData("key3", "Value from func Handler3(ctx iris.Context)")
	_, _ = ctx.JSON(ctx.GetViewData())
}
