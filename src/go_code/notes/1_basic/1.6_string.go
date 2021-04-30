package main

import "fmt"

func main() {
	my_string()
	// my_string()
}

func my_string_2() {
	/*
		1、字符串是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。
		   Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本

		2、Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本，这样 Golang 统一使用 UTF-8 编码
		   乱码问题就不会再困扰开发人员
		3、字符串一旦赋值了，字符串就不能修改，在 Go 中字符串是不可变的
		4、字符串的两种表示及形式
			1. 双引号，会识别转义字符
			2. 反引号，以字符串的原生形式输出，包括换行符和特殊字符，可以实现防止攻击
			   输出源代码等效果
	*/
	var address string = "中国 北京"
	fmt.Println("地址：", address)
	var str1 = `abc\nefg`
	fmt.Println("str1=", str1)
	var str2 = `
import "fmt"

func main() {
	my_string()
}
`
	fmt.Println(str2)

	// 字符串拼接
	var str3 string = "hello" + " world,"
	str3 += " hello linux"
	fmt.Println(str3)

	// 当一个拼接的操作很长时，可以分行写, 加号 "+" 必须放在上面
	var str4 = "hello world" +
		"hello linux" +
		"hello golang"
	fmt.Println(str4)
}

func my_string() {
	/*
		- 基本数据类型默认值
		整型		0
		浮点型		0
		字符串		""
		布尔类型	false
	*/
	var i int
	var f1 float32
	var f2 float64
	var b bool
	var s string
	fmt.Println("int 默认值:", i)
	fmt.Println("float32 默认值:", f1)
	fmt.Println("float64 默认值:", f2)
	fmt.Println("bool 默认值:", b)
	fmt.Println("string 默认值:", s)
}
