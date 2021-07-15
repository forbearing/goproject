package main

import "fmt"

// 多态数组演示
type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}
type Camera struct {
	Name string
}
type Computer struct{}

// 方法实现
func (this Phone) Start() {
	fmt.Println("Phone is start...")
}
func (this Phone) Stop() {
	fmt.Println("Phone is stop...")
}
func (this Camera) Start() {
	fmt.Println("Camera is start...")
}
func (this Camera) Stop() {
	fmt.Println("Camera is stop...")
}
func (this Computer) Work(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	// 定义一个 Usb 接口数组，可以存放 Phone 和 Camera 的结构体变量
	// 这里就体现出多态数组
	var usbArr [3]Usb
	fmt.Println(usbArr)
	usbArr[0] = Phone{"iPhone 12"}
	usbArr[1] = Phone{"Samsung"}
	usbArr[2] = Camera{"Nikon"}
	fmt.Println(usbArr)

}

/*
	1.基本介绍
		变量（实例）具有多种形态。面向对象的第三大特征，在 Go 语言，多态特征是通过接口实现的。
		可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态
	2. 接口体现多态特征
		1:多态参数
			在前面的 Usb 接口案例，Usb usb，既可以接收手机变量，又可以接收相机变量，就体现了 Usb 接口多态
		2:多态数组
			给 Usb 数组中，存放 Phone 就结构体和 Camara 结构体变量，Phone 还有一个特有的方法 call()
			请遍历 Usb 数组，如果是 Phone 变量，除了调用 Usb 接口声明的方法外，还需要调用 Phone 特有方法 call()
*/
