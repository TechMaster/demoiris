package controller

import (
	"demoiris/eris"
	"demoiris/logger"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/kataras/iris/v12"
)

/*
Hướng dẫn xử lý lỗi gọi REST API
*/
func (d demoError) CallAPI(ctx iris.Context) {
	if resp, err := http.Get("http://khongtontai.com/"); err != nil { //Kiểm tra lỗi kết nối
		logger.Log(ctx, eris.NewFrom(err))
	} else {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil { //Kiểm tra lỗi đọc dữ liệu
			ctx.WriteString(string(body))
		} else {
			logger.Log(ctx, eris.NewFrom(err))
		}
	}
}

func (d demoError) RetryCall(ctx iris.Context) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 3

	if resp, err := retryClient.Get("http://khongtontai.com/"); err != nil { //Kiểm tra lỗi kết nối
		logger.Log(ctx, eris.NewFrom(err))
	} else {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil { //Kiểm tra lỗi đọc dữ liệu
			ctx.WriteString(string(body))
		} else {
			logger.Log(ctx, eris.NewFrom(err))
		}
	}

}
