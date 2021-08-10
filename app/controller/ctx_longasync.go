package controller

import (
	"log"
	"time"

	"github.com/kataras/iris/v12"
)

func LongAsync(ctx iris.Context) {
	ctxCopy := ctx.Clone()
	go func() {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// note that you are using the copied context "ctxCopy", IMPORTANT
		log.Printf("Done! in path: %s", ctxCopy.Path())
	}()
	_, _ = ctx.WriteString("Trang trả về trước khi long task hoàn thành")
}
