package main

import "fmt"

// 定义一个 Cat 结构体，将 Cat 的各个字段信息，放入到 Cat 结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	var cat1 Cat
	fmt.Println("cat1 =", cat1)

	fmt.Println()
	var cat2 Cat
	cat2.Name = "小白"
	cat2.Age = 3
	cat2.Color = "白色"
	fmt.Println("猫的信息如下：")
	fmt.Println("name  =", cat2.Name)
	fmt.Println("age   =", cat2.Age)
	fmt.Println("color =", cat2.Color)
}

func notes() {
	/*
		===== 结构体和结构体变量（实例）的区别和联系 =====
		结构体
			1. 结构体是自定义数据类型，代表一类事物
			2. 结构体变量（实例）是具体的，代表一个具体变量
	*/
}
