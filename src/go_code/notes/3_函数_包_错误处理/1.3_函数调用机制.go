package main

import "fmt"

func main() {
	sum, sub := getSumAndSub(10, 5)
	fmt.Printf("两数之和 %d, 两数之差 %d\n", sum, sub)
	mySum, _ := getSumAndSub(20, 10)
	fmt.Printf("两数之和 %d\n", mySum)

}

func notes() {
	/*
		***** return 语句 *****
		1. Go 函数支持返回多个值，这一点是其它编程语言没有的
		2. func 函数名(形参列表) (返回值类型列表) {
			return 返回值列表
		}
			1. 如果返回多个值时，在接受时，希望忽略某个返回值，则使用 _ 符号表示占位忽略
			2. 如果返回值只有一个，（返回值类型列表）可以不写 ()
	*/
}

func getSumAndSub(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}
