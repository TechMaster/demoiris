package controller

import (
	"demoiris/logger"

	"demoiris/eris"

	"github.com/kataras/iris/v12"
)

func (d demoError) SerialCall(ctx iris.Context) {
	if err := funcA(); err != nil {
		logger.Log(ctx, err)
	} else {
		ctx.JSON("Success")
	}
}
func funcA() error {
	return funcB()
}

func funcB() error {
	return funcC()
}
func funcC() error {
	return eris.New("Cannot authenticate with OAuth2").EnableJSON()
}

func (d demoError) PanicCall(ctx iris.Context) {
	logger.Log(ctx, eris.Panic("Hard Disk is full"))
}
