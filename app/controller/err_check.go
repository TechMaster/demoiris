package controller

import (
	"demoiris/logger"
	"errors"

	"github.com/kataras/iris/v12"
)

func (d demoError) DemoCheckError(ctx iris.Context) {
	logger.CheckErr(ctx, funcE())
}

func funcE() error {
	return errors.New("This is simple error")
}
