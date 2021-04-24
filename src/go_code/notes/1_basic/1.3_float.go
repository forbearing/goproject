package main

import "fmt"

func main() {
	my_float()
}

func my_float() {
	// 浮点类型分类：单精度 float32，双精度: float64

	var price float32 = 91.23
	fmt.Println("price:", price)

	// 尾数部分可能丢失，造成精度损失
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1=", num1, "num2=", num2)

	// 浮点型使用细节
	// 1、Golang 浮点类型有固定的范围和字段长度，不受具体操作系统的影响。
	// 2、Golang 的浮点型默认声明为 float64 类型。
	// 3、浮点型常量有两种表示形式：十进制形式、科学计数法形式
	// 4、通常情况下，应该使用 float64，因为它比 float32 更精确，开发中，推荐使用 float64
	num3 := .111
	fmt.Printf("num3 数据类型是 %T", num3)
	num4 := 5.12345e2
	num5 := 5.12345e-2
	fmt.Println("num4=", num4, "num5=", num5)
}
