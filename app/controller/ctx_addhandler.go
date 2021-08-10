package controller

import "github.com/kataras/iris/v12"

/*
Thứ tự sẽ gọi lần lượt như sau
DemoAddHandler -> Handler4 -> adHocHandler
*/
func DemoAddHandler(ctx iris.Context) {
	ctx.ViewData("key1", "DemoAddHandler")
	ctx.AddHandler(adHocHandler) //Bổ xung handler này vào cuối cùng của handler stack
	ctx.Next()
}
func Handler4(ctx iris.Context) {
	ctx.ViewData("key4", "Handler4")
	ctx.Next()
}

func adHocHandler(ctx iris.Context) {
	ctx.ViewData("keyFinal", "adHocHandler")
	ctx.JSON(ctx.GetViewData())
}
