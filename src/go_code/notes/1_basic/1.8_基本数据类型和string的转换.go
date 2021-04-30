package main

import (
	"fmt"
	"strconv"
)

func main() {
	// baseic_to_string_1()
	// baseic_to_string_2()
	string_to_basic()
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
		在将 String 类型转成基本数据类型时，要确保 String 类型能够转成有效的数据，比如
		可以吧 "123" 转成一个整数，但不能把 "hello" 转成一个整数，如果这样做，Golang
		直接将其转成0

	*/
	var str1 string = "true"
	var str2 string = "1234"
	var str3 string = "9.87"
	var b bool
	var i int64
	var i2 int
	var f float64
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type %T, b value %v\n", b, b)
	i, _ = strconv.ParseInt(str2, 10, 64)
	i2, _ = strconv.Atoi(str2)
	fmt.Printf("i type %T, i value %v\n", i, i)
	fmt.Printf("i type %T, i value %v\n", i2, i2)
	f, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f type %T, f value %v\n", f, f)
}
