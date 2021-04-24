package main

import "fmt"

func main() {
	f1()
}

func f1() {
	/*
		1、Golang 和 C/Java 不同，Go 在不同类型的变量之间赋值时需要显式转换，也就是说 Golang
		   中数据类型不能自动转换
	*/

	var i int = 10
	var f float64 = float64(i)
	var u uint8 = uint8(f)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", u)
}
