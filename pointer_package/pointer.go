package main

import "fmt"


type person struct {
	name string
	age  int
}

func (n *person) changeInformation(informationObj person) {
	fmt.Print(informationObj)
}

func AnyPointerFunc() {
	// * dùng để khai báo kiểu dữ liệu con trỏ.
	// & dùng để lấy địa chỉ của một biến
	// * liên quan gì với &? trả lời giá trị của con trỏ là địa chỉ của biến mà con trỏ đó chỉ tới


	mrS := person{
		name: "Tung",
		age:  24,
	}

	updateInfo := person{
		name: "Tung2",
		age:  19,
	}

	mrS.changeInformation(updateInfo)
}