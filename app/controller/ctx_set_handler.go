package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

/*
Reset toàn bộ chuỗi handler v3.Get("/addhandler", controller.SetHandler, controller.Handler1, controller.Handler2, controller.Handler3)
về chỉ còn duy nhất adHocHandler
*/
func SetHandler(ctx iris.Context) {
	ctx.SetHandlers(context.Handlers{nil, Handler1, Handler3}) //Bỏ qua cái đầu tiên 0, giờ go next
	ctx.HandlerIndex(0)
	ctx.Next()
}
