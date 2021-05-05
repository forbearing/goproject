package main

import (
	"fmt"
)

func main() {
	// my_func1()
	// example1()
	example2()
}

func notes() {
	/*
		1. continue 语句用于结束本次循环，继续执行下一次循环，
		2. continue 语句出现在多层嵌套的循环语句体中时，可以通过标签指明要跳过的是哪一层循环，
	*/
}

func my_func1() {
	// label1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				continue
				// continue label1
			}
			fmt.Println("j=", j)
		}
	}
}
func example1() {
	// continue 实现 1-100 内的奇数，要求使用 for 循环 + continue
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
func example2() {
	// 从键盘输入个数不确定的整数，并判断读入的正数和负数的个数，输入为0时结束程序
	var positive int = 0
	var negetive int = 0
	for i := 0; ; i++ {
		var tempNum int
		fmt.Println("请输入一个数字")
		fmt.Scanln(&tempNum)
		if tempNum == 0 {
			break
		}
		if tempNum > 0 {
			positive++
			continue
		}
		negetive++
	}
	fmt.Println("结束")
	fmt.Println("正数个数:", positive)
	fmt.Println("负数个数:", negetive)
}
