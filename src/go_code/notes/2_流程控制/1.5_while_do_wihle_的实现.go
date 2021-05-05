package main

import "fmt"

func main() {
	my_func1()
}

func notes() {
	/*
		1. Go 没有 while 和 do while 语法，可以使用 for 循环来实现其使用效果。
		2. for 循环实现 while 的效果
			循环变量初始化
			for {
				if 循环条件表达式{break}
				循环体
				循环变量迭代
			}
		3. for 循环实现 do while 的效果
		   循环变量初始化
		   for {
			   循环体
			   循环变量迭代
			   if 循环条件表达式{break}
		   }
		   do-while 是先执行后判断，所以至少执行一次
	*/
}

func my_func1() {
	// while 循环实现
	var i int = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello golang")
		i++
	}

	// do-while 循环实现
	var j int = 1
	for {
		fmt.Println("hello linux")
		j++
		if j > 10 {
			break
		}
	}
}
