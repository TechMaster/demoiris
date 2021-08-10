package controller

import (
	"demoiris/logger"

	"demoiris/eris"

	"github.com/kataras/iris/v12"
)

//Trả về lỗi cùng thông tin bổ xung
func (d demoError) ErrWithData(ctx iris.Context) {
	if err := funcD(); err != nil {
		if errEris, ok := err.(*eris.Error); ok {
			data := map[string](interface{}){
				"host": "192.168.1.1",
				"port": 5432,
				"db":   "shopdb",
			}
			logger.Log(ctx, errEris.SetData(data))
		} else {
			logger.Log(ctx, err)
		}

	} else {
		ctx.JSON("Success")
	}
}

func funcD() error {
	return eris.SysError("Failed to connect to Postgresql")
}
