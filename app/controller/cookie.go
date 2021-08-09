package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

func Counter(ctx iris.Context) {
		action := ctx.Params().GetStringDefault("action", "")
	var counter int
	if action != "clear" {
		countVal := ctx.GetCookie("counter")
		if c, err := strconv.Atoi(countVal); err == nil {
			counter = c + 1
		} else {
			counter = 0
		}

	} else if action == "clear" {
		counter = 0
	}

	ctx.SetCookie(&iris.Cookie{
		Name:    "counter",
		Value:   fmt.Sprint(counter),
		Expires: time.Now().Add(48 * time.Hour),
	})

	ctx.ViewLayout("default")
	_ = ctx.View("counter", iris.Map{
		"counter": counter,
	})
}
