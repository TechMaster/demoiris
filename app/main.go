package main

import (
	"demoiris/logger"
	"demoiris/router"
	"demoiris/template"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New() //chỉ tạo một iris app blank không dùng middleware mặc định

	logFile := logger.InitErrLog("logs/")
	defer logFile.Close()

	app.HandleDir("/", "../static")
	app.UseRouter(recover.New())
	//app.Use(middleware.LogMiddleware)

	router.RegisterRoutes(app)
	template.InitViewEngine(app)

	if err := app.Listen(":8080"); err != nil {
		app.Logger().Fatal("Failed to start app")
	}
}
