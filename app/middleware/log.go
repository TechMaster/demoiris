package middleware

import "github.com/kataras/iris/v12"

func LogMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
