package controller

import (
	"demoiris/eris"
	"demoiris/logger"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris/v12"
)

func (d demoError) ShowEmailForm(ctx iris.Context) {
	_ = ctx.View("email")
}

func (d demoError) ProcessEmailForm(ctx iris.Context) {
	email := ctx.PostValue("email")
	if !govalidator.IsEmail(email) {
		_, _ = ctx.JSON("Email is invalid")
	} else {
		_, _ = ctx.JSON("Success")
	}
}

func (d demoError) GetWarning(ctx iris.Context) {
	logger.Log(ctx, eris.Warning("Your email is invalid").EnableJSON().StatusCode(iris.StatusBadRequest))
}
