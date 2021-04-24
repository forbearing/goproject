package main

import "fmt"

// 声明多个全局变量
var g1 = 100
var g2 = 200
var gname = "hybfkuf"

// 一次性声明多个全局变量
var (
	g3     = 100
	g4     = 200
	gname2 = "hybfkuf2"
)

func main() {
	//how_to_define_var()
	//output_global_var()
	//precautions_for_use()
	plus_sign()
}

// 三种定义变量的方法
func how_to_define_var() {
	// 第一种：指定变量类型，声明后若不赋值，使用默认值
	var i int
	fmt.Println("i=", i)

	// 第二种：根据值自行判断变量类型（类型推导）
	var num = 10 // var num int 10
	var num2 = 10.13
	fmt.Println("num=", num)
	fmt.Println("num2=", num2)

	// 第三种：省略 var，注意 := 左侧变量不应该是已经声明过的，否则会导致编译错误
	var1 := "string1"
	var var2 string = "string2"

	fmt.Println("var1=", var1)
	fmt.Println("var2=", var2)

	// 第四种：多变量声明，在编程中，一次性声明多个变量
	var n1, n2, n3 int // 同时声明多个变量
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	var n4, name, n5 = 100, "tom", 88
	fmt.Println("n4=", n4, "name=", name, "n5=", n5)
	n6, name2, n7 := 100, "Jonas", 10
	fmt.Println("n6=", n6, "name=", name2, "n7=", n7)
}

// 输出全局变量
func output_global_var() {
	fmt.Println("g1=", g1, "g2=", g2, "gname=", gname)
	fmt.Println("g3=", g3, "g4=", g4, "gname2=", gname2)
}

// 变量使用的注意事项
func precautions_for_use() {
	// 1、该区域的数据值可以在同一类型范围内不断变化
	var i int = 10
	i = 30
	i = 50
	//i = 10.1 // 变量 i 原来是整形，不能变成浮点型，不能改变数据类型。
	fmt.Println("i=", i)

	// 2、变量在同一作用域（在一个函数或者代码块）内不能重名
	// 3、变量 = 变量名+值+数据类型（变量三要素）
	// 4、Golang 的变量如果没有赋初值，编译器会使用默认值，比如 int 默认值为0，string 默认值为空
}

// 加号的使用
func plus_sign() {
	// 当左右两边都是数值型时，则做加法运算
	// 当左右两边都是字符串，则做字符串拼接
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Print("sum=", sum)
	str1 := "hello "
	str2 := "world"
	mystr := str1 + str2
	fmt.Println("my string is", mystr)
}
