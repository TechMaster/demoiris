package router

import (
	"demoiris/controller"

	"github.com/kataras/iris/v12"
)

func RegisterRoutes(app *iris.Application) {
	app.Get("/", controller.ShowHomePage)
	route_error(app)
}

func route_error(app *iris.Application) {
	v1 := app.Party("/err")
	{
		v1.Get("/serialcall", controller.Err.SerialCall)
		v1.Get("/panic", controller.Err.PanicCall)
		v1.Get("/errdata", controller.Err.ErrWithData)
		v1.Get("/checkerr", controller.Err.DemoCheckError)
		v1.Get("/inputmail", controller.Err.ShowEmailForm)
		v1.Post("/submitemail", controller.Err.ProcessEmailForm)
		v1.Get("/warning", controller.Err.GetWarning)
		v1.Get("/callapi", controller.Err.CallAPI)
		v1.Get("/retrycall", controller.Err.RetryCall)
	}

	v2 := app.Party("/cooksess")
	{
		v2.Get("/counter/{action}", controller.Counter)
	}

	v3 := app.Party("/ctx")
	{
		v3.Get("/checkctxpage", controller.ShowCheckCtx)
		v3.Get("/checkctx", controller.CheckTypeOfCtx)
		v3.Get("/normalget", controller.CheckTypeOfCtx)
		v3.Get("/ajaxget", controller.CheckTypeOfCtx)
		v3.Get("/ajaxpost", controller.CheckTypeOfCtx)
		v3.Get("/redirect", controller.RedirectCheckCtx)
		v3.Get("/chainhandlers", controller.Handler1, controller.Handler2, controller.Handler3)
		v3.Get("/render_markdown", controller.RenderMarkDown)
		v3.Get("/addhandler", controller.DemoAddHandler, controller.Handler4)
		v3.Get("/sethandler", controller.SetHandler, controller.Handler1, controller.Handler2, controller.Handler3)
		v3.Get("/getstopwitherror", controller.GetStopWithError)
		v3.Post("/poststopwitherror", controller.PostStopWithError)
		v3.Get("/longasync", controller.LongAsync)
	}
}
