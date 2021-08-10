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
	if govalidator.IsEmail(email) {
		_, _ = ctx.JSON("Email is valid")
	} else {
		logger.Log(ctx, eris.Warning("Email is invalid").StatusCode(iris.StatusBadRequest))
	}
}

func (d demoError) GetWarning(ctx iris.Context) {
	logger.Log(ctx, eris.Warning("Your email is invalid").EnableJSON().StatusCode(iris.StatusBadRequest))
}
