package template

import "github.com/kataras/iris/v12"

func InitViewEngine(app *iris.Application) {
	viewEngine := iris.Blocks("./views", ".html")
	viewEngine.Layout("default")

	
	app.RegisterView(viewEngine)
}
