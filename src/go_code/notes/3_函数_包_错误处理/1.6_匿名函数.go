package main

import "fmt"

// 全局匿名函数
// 变量名必须大写才能全局有效
var (
	Sum = func(n1 int, n2 int) int {
		return n1 + n2
	}
	Sub = func(n1 int, n2 int) int {
		return n1 - n2
	}
)

func main() {

	// 用匿名函数完成两个数的求和
	res := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println(res)

	// 将匿名函数赋值给一个变量
	func1 := func(num1 int, num2 int) int {
		return num1 + num2
	}
	res1 := func1(10, 30)
	fmt.Println(res1)

	// 全局匿名函数的使用
	res2 := Sum(4, 9)
	fmt.Println(res2)
	res3 := Sub(9, 4)
	fmt.Println(res3)

}

func notes() {
	/*
		***** 匿名函数 *****
		1. Go 支持匿名函数，如果我们某个函数只希望使用一次，可以考虑使用匿名函数，匿名函数也可以实习多次调用
		2. 匿名函数的使用方式1
			1. 在定义匿名函数的时候就直接调用
			2. 将匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数
		3. 如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以在程序有效。
	*/
}
