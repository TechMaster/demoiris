package controller

import (
	"demoiris/eris"
	"demoiris/logger"

	"github.com/goccy/go-json"
	"github.com/kataras/iris/v12"
)

func ShowCheckCtx(ctx iris.Context) {
	ctx.ViewLayout("default")
	ctx.View("checkctx")
}
func CheckTypeOfCtx(ctx iris.Context) {
	data := iris.Map{
		"isAJAX":         ctx.IsAjax(),
		"isMobile":       ctx.IsMobile(),
		"isSSL":          ctx.IsSSL(),
		"isDebug":        ctx.IsDebug(),
		"isHTTP2":        ctx.IsHTTP2(),
		"isScript":       ctx.IsScript(),
		"RouteName":      ctx.RouteName(),
		"FullRequestURI": ctx.FullRequestURI(),
		"Host":           ctx.Host(),
		"ContentType":    ctx.GetContentType(),
		"Referrer":       ctx.GetReferrer(), //redirect
		"HandlerName":    ctx.HandlerName(),
		"StatusCode":     ctx.GetStatusCode(),
	}
	if ctx.IsAjax() {
		ctx.JSON(data)
	} else {
		if dataStr, err := json.Marshal(data); err == nil {
			ctx.WriteString(string(dataStr))
		} else {
			logger.Log(ctx, eris.NewFrom(err))
		}

	}
}
func RedirectCheckCtx(ctx iris.Context) {
	ctx.Redirect("/ctx/checkctx")
}
