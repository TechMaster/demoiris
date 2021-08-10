package controller

import (
	"demoiris/eris"
	"demoiris/logger"
	"io/ioutil"

	"github.com/kataras/iris/v12"
)

func RenderMarkDown(ctx iris.Context) {
	if bytes, err := ioutil.ReadFile("../ReadMe.md"); err != nil {
		logger.Log(ctx, eris.NewFromMsg(err, "Cannot open ReadMe.md"))
	} else {
		ctx.Markdown(bytes)
	}

}
