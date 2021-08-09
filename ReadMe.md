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