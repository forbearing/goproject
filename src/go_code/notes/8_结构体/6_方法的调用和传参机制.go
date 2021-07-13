package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func main() {
	// 声明一个结构体 Circle，字段为 radius
	// 声明一个方法 area 和 Circle 绑定，可以返回面积
	c := Circle{4.0}
	res := c.area()
	fmt.Println(res)
}

func notes() {
	/*
		===== 方法的调用和传参机制原理 ======
		1. 在通过一个变量去调用方法时，其调用机制和函数一样
		2. 不一样的地方是，变量调用方法时，该变量本身也会作为一个参数传递到方法（如果变量是值类型，则进行值拷贝，
		   如果变量是引用类型，则进行地址拷贝）。
	*/
}
