package main

import "fmt"

func main() {
	description()

}

func description() {
	/*
		运算符种类
			1. 算术运算符
			2. 赋值运算符
			3. 比较运算符/关系运算符
			4. 逻辑运算符
			5. 位运算符
			6. 其他运算符 & *（Go 没有三元运算符）

		1. 算数运算符
			1. 对于除号 "/" ,它的整数和除是有区别的：整数除之间做除法，只保留整数部分而丢弃
			   小数部分，例如：x := 19/5，结果是3
			2. 当对一个数取模时，可以等价于 a%b = a-a / b*b，这样就可以看作是一个取模的本质运算
			3. Golang 的自增自减只能当作一个独立语言使用，不能这样使用 b:= a++ 或者 b := a--
			4. Golang 的 ++ 和 -- 只能写在变量的后面，不能写在前面，即：只有 a++ a--，没有 ++a --a
			5. Golang 的设计者去掉 C/Java 中的自增自减的容易混淆的写法，让 Golang 更加简洁，统一（强制性的）



	*/

	// 算数运算符
	var num1 float32 = 10 / 4
	fmt.Println(10 / 4)
	fmt.Println(num1)

	var i = 10
	fmt.Println(i)
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
}
