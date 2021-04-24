package main

import "fmt"

func main() {
	f1()
}

func f1() {
	/*
		1、Golang 和 C/Java 不同，Go 在不同类型的变量之间赋值时需要显式转换，也就是说 Golang
		   中数据类型不能自动转换
		2、Go 中，数据类型的转换可以是从 “表示范围小 -> 表示范围大”，也可以反过来
		3、被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
		4、在转换中，比如将 int64 转换成 int8，编译时不会报错，只是转换的结果是按溢出处理，
		   和我们喜欢的结果不一样。
	*/

	var i int = 10
	var f float64 = float64(i)
	var u uint8 = uint8(f)
	var i2 int64 = int64(i)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", i2)
}
