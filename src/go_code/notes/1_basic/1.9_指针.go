package main

import "fmt"

func main() {
	// my_point()
	case_1()
}

func my_point() {
	/*
		1、基本介绍
			1、基本数据类型，变量存的就是值，也叫值类型
			2、获取变量的地址，用 &、比如：var num int，获取 num 的地址： &num
			3、指针类型，变量存的是一个地址，这个地址指向的空间才是值，比如：var ptr *int = &num
			4、获取指针类型所指向的值，使用：* 比如 var *ptr int, 使用 *ptr 获取 p 指向的值。
		2、指针使用细节
			1、值类型，都有对应的指针类型，形式为 *数据类型，比如 int 的对应的指针就是 *int
				float32 对应的指针类型就是 *float，依次类推
			2、值类型包括：基本数据类型 int 系列、float 系列，bool、string、数组和结构体 struct
	*/
	var i int = 10
	fmt.Printf("i 的地址 %v\n", &i)
	var p *int = &i
	fmt.Println("p 的值:", p)
	fmt.Println("p 的地址:", &p)
	fmt.Println("p 指向地址的值:", *p)
}

func case_1() {
	// 通过指针来修改变量的值
	var num int = 99
	fmt.Println("num 的值:", num)
	var ptr *int = &num
	*ptr = 100
	fmt.Println("num 的值:", num)
}
