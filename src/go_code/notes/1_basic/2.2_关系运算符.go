package main

import "fmt"

func main() {
	my_func()
}

func my_func() {
	/*
		1. 介绍
			1. 关系运算符的结果都是 bool 类型，也就是要么是 true，要么是 false。
			2. 关系表达式经常用在 if 结构的条件中或循环结构的条件中。
		2. 细节说明
			1. 关系运算符组成的表达式，我们称为 “关系表达式”：a > b
			2. 比较运算符 "==" 不能写成 赋值运算符 "="
	*/
	var n1 = 9
	var n2 = 8
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 <= n2)
}
