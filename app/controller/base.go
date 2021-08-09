package controller

/*
Tạo ra một cấu trúc rỗng rồi sau đó định nghĩa hàm
func (d demoError) DemoCheckError(ctx iris.Context) {
	logger.CheckErr(ctx, funcE())
}
Khi gọi sẽ có đối tượng Err đứng controller để phân biệt rõ hàm này thuộc nhóm nào
v1.Get("/checkerr", controller.Err.DemoCheckError)
*/
var Err = demoError{}

type demoError struct {
	//Sau này cần có thể nhồi dữ liệu vào đây. Các hàm gắn với đối tượng này sẽ truy cập được dễ dàng
}
