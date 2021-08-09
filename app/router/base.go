package router

import (
	"demoiris/controller"
	"demoiris/logger"

	"github.com/kataras/iris/v12"
)

func RegisterRoutes(app *iris.Application) {
	app.Get("/", showHomePage)
	route_error(app)
}

func showHomePage(ctx iris.Context) {
	ctx.ViewLayout("default")
	viewData := iris.Map{ //iris.Map tương đương map[string]interface{}
		"Title":   "Siêu to khổng lồ",              //--> layouts/default.html -> {{.Title}}
		"address": "Tầng 12A, Viwaseen, 48 Tố Hữu", //--> partials/footer.html --> {{.address}}
		"UserInfo": iris.Map{ //--> partials/menu.html -->
			"Name":  "Cường",
			"Email": "cuong@techmaster.vn",
		},
	}
	logger.CheckErr(ctx, ctx.View("index", viewData))
}

func route_error(app *iris.Application) {
	v1 := app.Party("/err")
	{
		v1.Get("/serialcall", controller.Err.SerialCall)
		v1.Get("/panic", controller.Err.PanicCall)
		v1.Get("/jsonerr", controller.Err.JSONErr)
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
}
