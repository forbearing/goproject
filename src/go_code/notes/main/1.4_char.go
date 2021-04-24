package main

import "fmt"

func main() {
	//my_char1()
	my_char2()
}

// 1、字符使用
func my_char1() {
	// 1、Golang 中没有专门的字符类型，如果要存储单个字符（字母），一般使用 byte 来保存
	// 2、字符串是一串固定长度的字符连接起来的字符序列，而 Go 的字符串不同，它是由字符组成的。
	// 3、而 Go 的字符串是由单个字节连接起来的，也就是说对于传统的字符串是由字符组成的。
	// 4、如果我们保存的字符在 ASCII 表中，可以直接保存到 byte
	//	  如果我们保存的字符对应码值大于 255，这时我们可以考虑使用 int 类型。
	// 5、如果我们需要按照字符的方式输出，此时我们需要格式化输出 fmt.Printf()

	// 当我们直接输出 byte 值，就是输出了对应字符的 ASCII 码值
	// 如果希望输出字符，需要使用格式化输出
	var c1 byte = 'a'
	var c2 byte = '0'
	//var c3 byte = '字'  overflows 溢出
	var c3 int = '字'
	fmt.Println("c1=", c1, "c2=", c2)
	fmt.Printf("c1= %c c2= %c\n", c1, c2)
	fmt.Printf("c3= %c", c3)
}

// 2、字符使用细节
func my_char2() {
	/*
		字符使用细节
		1、字符常量是用单引号括起来的字符，比如:
			var c1 byte ='a',
			var c2 int = '字',
			var c3 byte = '9'
		2、Go 中允许使用转义字符 '\' 来将其后的字符转变成特殊字符常量，例如：
			var c3 char = '\n'		// \n 表示换行
		3、Go 语言的字符使用 UTF-8 编码
		4、在 Go 中，字符的本质是一个整数，直接输出时，是该字符对应的 UTF-8 编码的码值
		5、可以直接给某个变量赋一个数字，然后按格式化输出时 %c 会输出该数字对应的 unicode 字符
		6、字符类型是可以进行运算的，相当于一个整数，因为它都对应有 Unicode 码
	*/
	/*
		字符类型本质讨论
		1、字符型存储到计算机，需要将字符对应的码值（整数）找出来
			存储：字符 --> 对应码值 --> 二进制 --> 存储
			读取：二进制 --> 码值 --> 字符 --> 读取
		2、字符和码值的对应关系是通过字符编码来决定的（是规定好的）
		3、Go 语言的编码都统一成了 utf-8，和其他编程语言来说，非常方便，很统一，再也没有编码乱码的困扰了。
	*/
	var c1 int = 22269
	fmt.Printf("c1=%c\n", c1)
	var n1 = 10 + 'a'

	// 字符串进行运算，相当于一个整数，运算时是按照码值进行
	fmt.Println("n1=", n1)
}
