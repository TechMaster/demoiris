package controller

import (
	"github.com/kataras/iris/v12"
)

func GetStopWithError(ctx iris.Context) {
	ctx.ViewLayout("default")
	_ = ctx.View("stopwitherror")
}

//Iris cung cấp sẵn hàm trả về lỗi ctx.StopWithError
func PostStopWithError(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	data := iris.Map{
		"id":      id,
		"page":    ctx.URLParamIntDefault("page", 0),
		"name":    ctx.PostValue("name"),
		"message": ctx.PostValue("message"),
	}
	_, _ = ctx.JSON(data)

}
