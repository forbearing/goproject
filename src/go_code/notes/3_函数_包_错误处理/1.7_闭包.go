package main

import (
	"fmt"
	"strings"
)

func main() {
	// func1 := AddUpper()
	// fmt.Println(func1(1))
	// fmt.Println(func1(2))
	// fmt.Println(func1(3))

	suffixConf := makeSuffix(".conf")
	name := suffixConf("test.txt")
	fmt.Println(name)
	suffixExe := makeSuffix(".exe")
	name = suffixExe("test.txt")
	fmt.Println(name)

}
func notes() {
	/*
		***** 介绍 *****
		1. 闭包就是一个函数和其相关的引用环境组合的一个整体（实体）

		***** 闭包的最佳实践 *****
		编写一个程序，具体要求如下
		1. 编写一个函数，makeSuffix(suffix string) 可以接收一个文件后缀名(比如 .jpg), 并返回一个闭包
		2. 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如 .jpg)，则返回文件名.jpg
		   如果已经有 .jpg 后缀，则返回原文件名
		3. 要求使用闭包方式完成
		4. strings.HasSuffix
	*/
}

func my_func1() {
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 累加器，返回一个函数
func AddUpper() func(int) int {
	var n int = 10
	// 1. 返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包。
	// 2. 可以理解成：闭包是类，函数是操作，n是字段。函数和它使用到的n构成闭包。
	// 3. 当 main 函数中，反复调用 func1 函数，因为 n 只初始化一次，因此每次调用一次就是进行累计
	// 4. 搞清楚闭包的关键，就是要分析出返回的函数它使用(引用)到哪些变量。因为函数和它引用到的变量共同构成闭包。
	return func(x int) int {
		n = n + x
		return n
	}
}
