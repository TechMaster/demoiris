# Demo các tính năng chính của Iris

### Tạo lỗi, báo lỗi, log lỗi

Mã nguồn của eris để luôn ở thư mục [eris](app/eris)

Logic log lỗi để ở [logger](app/logger)

Cấu hình log lỗi ra file
```go
logFile := logger.InitErrLog("logs/")
defer logFile.Close()
```
File logs được lưu ở thư mục [logs](app/logs)

Trong thư mục [controller](app/controller) có rất nhiều ví dụ log lỗi
```
.
├── base.go
├── err_call_rest.go
├── err_check.go
├── err_json.go
├── err_serial_call.go
└── err_warning.go
```

Do các hàm handler của iris không trả về error do đó nếu có lỗi ở đâu là phải xử lý luôn
```go
func (d demoError) SerialCall(ctx iris.Context) {
	if err := funcA(); err != nil {
		logger.Log(ctx, err)  //Xử lý lỗi luôn
	} else {
		ctx.JSON("Success")
	}
}
```

Chi tiết hơn xem
[https://github.com/TechMaster/eris](https://github.com/TechMaster/eris)


## Gọi REST API đúng cách

### Dùng http client chuẩn
```go
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
```

### Dùng retry http client
Dùng [go-retryablehttp](github.com/hashicorp/go-retryablehttp)
```go
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
```

## View Template Engine

Xem [template/base.go](app/template/base.go)

Cấu hình view template engine, Blocks (nâng cấp của thư viện HTML)
```go
func InitViewEngine(app *iris.Application) {
	viewEngine := iris.Blocks("./views", ".html")
	viewEngine.Layout("default")
	app.RegisterView(viewEngine)
}
```
Cấu trúc thư mực view template
```
view
├── layouts  --> Các layout template để chứa các template khác
│   └── default.html
├── partials  --> Chứa các một thành phần tái sử dụng của template
│   ├── footer.html
│   ├── header.html
│   └── menu.html
├── counter.html
├── email.html
├── error.html
└── index.html
```

Đổ dữ liệu vào ViewTemplate
```go
func showHomePage(ctx iris.Context) {
	ctx.ViewLayout("default")
	viewData := iris.Map{ //iris.Map tương đương map[string]interface{}
		"Title":   "Siêu to khổng lồ",              //--> layouts/default.html -> {{.Title}}
		"address": "Tầng 12A, Viwaseen, 48 Tố Hữu", //--> partials/footer.html --> {{.address}}
		"UserInfo": iris.Map{ //--> partials/menu.html -->
			"Name":  "Cường",
			"Email": "cuong@techmaster.vn",
		},
	}
	logger.CheckErr(ctx, ctx.View("index", viewData))
}
```