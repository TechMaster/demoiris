package controller

import (
	"demoiris/email"
	"demoiris/logger"

	"github.com/kataras/iris/v12"
)

func SendSimpleEmail(ctx iris.Context) {
	sendGrid := email.SendGrid{}
	if err := sendGrid.SendSimple("cuong@techmaster.vn", "I love you Cuong", "Today is good weather"); err != nil {
		logger.Log(ctx, err)
	} else {
		ctx.WriteString("Send success")
	}
}
