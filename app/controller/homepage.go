package controller

import (
	"demoiris/logger"

	"github.com/kataras/iris/v12"
)

func ShowHomePage(ctx iris.Context) {
	ctx.ViewLayout("default")
	viewData := iris.Map{ //iris.Map tương đương map[string]interface{}
		"Title":   "Demo Iris V12 core features",              //--> layouts/default.html -> {{.Title}}
		"address": "Tầng 12A, Viwaseen, 48 Tố Hữu", //--> partials/footer.html --> {{.address}}
		"UserInfo": iris.Map{ //--> partials/menu.html -->
			"Name":  "Cường",
			"Email": "cuong@techmaster.vn",
		},
	}
	logger.CheckErr(ctx, ctx.View("index", viewData))

}
