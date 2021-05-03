package main

import "fmt"

func main() {
	// my_func()
	my_func2()
}

func my_func() {
	/*
		fmt.Scanln()		# 获取一行的输入，在换行时停止扫描，最后一个条目后必须有换行或者到达结束位置。
		fmt.Scanf()			# 从标准输入扫描文本，根据 format 参数指定的格式成功读取的空白分隔符。
							# 的值保存进成功传递给本函数的参数。返回成功扫描的条目个数和遇到的任何错误。
	*/

	// 方式一：fmt.Scanln()
	var name string
	var age byte
	var salary float32
	var isPass bool
	fmt.Println("请输入姓名: ")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄: ")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水: ")
	fmt.Scanln(&salary)
	fmt.Println("请输入是否通过考试: ")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是: %v\n年龄是 %v\n薪水是 %v\n是否通过考试 %v\n", name, age, salary, isPass)
}

func my_func2() {
	var name string
	var age byte
	var salary float32
	var isPass bool

	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开:")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	fmt.Printf("名字是: %v\n年龄是 %v\n薪水是 %v\n是否通过考试 %v\n", name, age, salary, isPass)
}
