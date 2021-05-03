package main

import "fmt"

func main() {
	my_func()
}

func my_func() {
	/*
		1. 在 golang 中不能直接使用二进制来表示一个整数，它沿用了C的语法特点
		2. 八进制，以数字0开头表示
		3. 十六进制, 以 0x 或 0X 开头表示
	*/

	var i int = 10
	fmt.Printf("10的二进制：%b\n", i)

	var j int = 011
	fmt.Printf("j %d\n", j)

	var k int = 0x11
	fmt.Printf("k %d\n", k)
}
