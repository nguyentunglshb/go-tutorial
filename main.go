package main

import (
	"fmt"
	"go-tutotial/testpackage"
)

func main() {
	var step int = 5
	testpackage.AnyFunc(step)

	// * dùng để khai báo kiểu dữ liệu con trỏ.
	// & dùng để lấy địa chỉ của một biến
	// * liên quan gì với &? trả lời giá trị của con trỏ là địa chỉ của biến mà con trỏ đó chỉ tới

	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
}
