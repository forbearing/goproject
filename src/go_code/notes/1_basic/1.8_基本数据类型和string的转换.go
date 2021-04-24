package main

import (
	"fmt"
	"strconv"
)

func main() {
	// baseic_to_string_1()
	// baseic_to_string_2()
}

func baseic_to_string_1() {
	/*
		基本类型转字符串
			方法一：fmt.Sprintf("%参数", 表达式)
				1、参数需要和表达式的数据类型相匹配
				2、fmt.Sprintf() 会返回转换后的字符串
			方法二：使用 strconv 包函数
			方法三：ItoA 函数
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

	var num1 int = 100
	var num2 int64 = 200
	str = strconv.Itoa(num1)
	str = strconv.Itoa(int(num2))
	fmt.Printf("str type %T, str value %q", str, str)
}

func baseic_to_string_2() {
	/*
		strconv.FormatInt()
		strconv.FormatFloat()
		strconv.FormatBool()
		strconv.FormatUnit()

		strconv.FormatFloat(f, 'f', 10, 64)
			'f'		格式
			10		小数位保留10位
			64		小数是 float64
	*/
	var i int = 99
	var f float64 = 23.456
	var b bool = true
	// var c byte = 'h'
	var str string

	str = strconv.FormatInt(int64(i), 10) // 10 表示十进制
	fmt.Printf("str type %T, str value %q\n", str, str)
	str = strconv.FormatFloat(f, 'f', 10, 64)
	fmt.Printf("str type %T, str value %q\n", str, str)
	str = strconv.FormatBool(b)
	fmt.Printf("str type %T, str value %q\n", str, str)
}

func string_to_basic() {
	/*
		字符串转基本类型：
			方法一：
				strconv.ParseBool()
				strconv.ParseInt()
				strconv.ParseFloat()
				strconv.ParseUnit()
			方法二：Atoi
		注意
			b, _ = strconv.ParseBool(str)
			ParseBool 会会返回两个值，上面表示只获取第一个值，第二个值忽略

	*/
	var str1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T, b value %q\n", b, b)
}
