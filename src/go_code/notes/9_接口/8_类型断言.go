package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point // ok
	var b Point
	// b = a // error
	// b = a.(Point) 就是类型断言，表示判断 a 是否指向 Point 类型的变量，如果是就转成 Point
	// 类型并赋值给 b 变量，否则报错。
	b = a.(Point)
	fmt.Println(b)

	var hello interface{}
	var i1 = 100
	hello = i1
	var i2 int = hello.(int)
	fmt.Println(hello)
	fmt.Println(i2)

	var mygo interface{}
	var b1 bool = true
	mygo = b1
	var b2 bool = mygo.(bool)
	fmt.Println(mygo)
	fmt.Println(b2)
	var linux interface{}
	var f1 float64 = 100.111
	linux = f1
	// 类型断言
	// if f2, ok := linux.(float64); ok {
	if f2, ok := linux.(float32); ok {
		fmt.Printf("f2 的类型是 %T, 值是 %v\n", f2, ok)
	} else {
		fmt.Println("类型断言失败")
	}
}

/*
	类型断言：由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
	例如：
		var t float32
		var x interface{}
		x = t
		y := x.(float32)

	1.在进行类型断言时，如果类型不匹配就回报 panic，因此进行类型断言时，要确保原来的空接口指向的就是要断言的类型。
	2.如何在进行断言时，带上检测机制，如果成功就 ok，失败也不要报 panic
*/
