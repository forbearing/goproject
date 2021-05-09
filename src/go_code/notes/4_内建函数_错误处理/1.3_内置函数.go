package main

import "fmt"

func main() {
	new_usage()
}

func notes() {
	/*
		len			求长度
		new			用来分配内存，主要用来分配值类型，比如 int、float32、struct... 返回的是指针
		make		用来分配内存，主要用来分配引用内存，比如 chan、map、slice
	*/
}

func new_usage() {
	num1 := 100
	fmt.Printf("num1的类型:%T, num1的值:%v, num1的地址%v\n", num1, num1, &num1)
	num2 := new(int)
	fmt.Printf("num2的类型:%T, num2的值:%v, num2的地址%v, num2指向的值:%v\n", num2, num2, &num2, *num2)
}
