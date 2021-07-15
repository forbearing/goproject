package main

import "fmt"

// 给 Phone 结构体增加一个特有的 call(), 当 Usb 接口接收的是 Phone 变量时，还需要调用 call() 方法
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
	fmt.Println(this.Name, "is start...")
}
func (this Phone) Stop() {
	fmt.Println(this.Name, "is stop...")
}
func (this Phone) Call() {
	fmt.Println(this.Name, "is calling...")
}
func (this Camera) Start() {
	fmt.Println(this.Name, "is start...")
}
func (this Camera) Stop() {
	fmt.Println(this.Name, "is stop...")
}
func (this Computer) Work(usb Usb) {
	usb.Start()
	// 如果 usb 是指向 Phone 结构体变量，则还需要调用 Call 方法
	// 类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var usbArr [3]Usb
	fmt.Println(usbArr)
	usbArr[0] = Phone{"iPhone 12"}
	usbArr[1] = Phone{"Samsung"}
	usbArr[2] = Camera{"Nikon"}
	fmt.Println(usbArr)
	fmt.Println()

	var computer Computer
	for _, v := range usbArr {
		computer.Work(v)
	}
}
