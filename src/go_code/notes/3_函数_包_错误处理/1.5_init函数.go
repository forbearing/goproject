package main

import (
	"fmt"
	"hybfkuf/utils"
)

var age = test()

func test() int {
	fmt.Println("test()...")
	return 90
}

func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()... age=", age)
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
}

func notes() {
	/*
		***** 基本介绍 *****
		1. 每一个源文件都可以包含一个 init 函数，该函数会在 main 函数执行前，被 Go 运行框架调用
		   也就是说 init 在 main 函数前调用
		2. 通常可以在 init 函数中完成初始化工作。

		***** 注意事项和细节 *****
		1. 如果一个文件同时包含全局变量定义，init 函数，main 函数，则执行流程是
		   变量定义 -> init 函数 -> main 函数
		2. 如果引入了其它包，则执行流程是
		   包中的变量定义 -> 包中的 init 函数 -> 本程序的变量定义 -> 本程序的 init 函数 -> 本程序的 main 函数
	*/
}
