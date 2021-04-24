package main

import "fmt"

func main() {
	f1()
}

func f1() {
	/*
		方法一：fmt.Sprintf("%参数", 表达式)
			1、参数需要和表达式的数据类型相匹配
			2、fmt.Sprintf() 会返回转换后的字符串
		方法二：使用 strconv 包函数
	*/
	var i int = 99
	var f float64 = 23.456
	var b bool = true
	var c byte = 'h'
	var str string
	str = fmt.Sprintf("%v", i)
	fmt.Printf("str type: %T, str value: %q\n", str, str)
	str = fmt.Sprintf("%f", f)
	fmt.Printf("str type: %T, str value: %q\n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type: %T, str value: %q\n", str, str)
	str = fmt.Sprintf("%c", c)
	fmt.Printf("str type: %T, str value: %q\n", str, str)
}
