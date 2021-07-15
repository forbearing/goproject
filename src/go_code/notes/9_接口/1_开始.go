package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Phone struct{}
type Camera struct{}
type Computer struct{}

// 让 Phone 实现 Usb 接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

// 让 Camera 实现 Usb 接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作 ...")
}

// 计算机
// 编写一个方法 Working 方法，接收一个 Usb 接口类型变量
// 只是实现了 Usb 接口，所谓实现 Usb 接口就是指实现了 Usb 接口声明的所有方法

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	// 先创建结构体变量
	phone := Phone{}
	camera := Camera{}
	computer := Computer{}

	// 使用接口
	computer.Working(phone)
	computer.Working(camera)
}

func notes() {
	/*
		接口：松耦合，高内聚 的编程思想
	*/
}
