package main

import "fmt"

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}

func notes() {
	/*
		***** 为什么需要 defer *****
		1. 在函数中，程序员经常需要创建资源，比如：数据库连接、文件句柄、锁等。为了在函数执行完毕后，及时的
		   释放资源，Go 的设计者提供 defer（延时机制）

		***** defer 细节说明 *****
		1. 当 go 执行到一个 defer 时，不会立即执行 defer 后的语句，而是将 defer 后的语句压入到一个
		   栈中，然后继续执行函数下一个语句	。
		2. 当函数执行完毕后，在从 defer 栈中，依次从栈顶取出语句执行，先入后出的机制。
		3. 在 defer 将语句放入到栈时，也会将相关的值拷贝同时入栈。

		***** defer 最佳实践 ******
		1. defer 最主要的价值是在，当函数执行完毕后，可以及时的释放函数创建的资源。
			// 关闭文件资源
			file = openfile(文件名)
			defer file.close()
			// 释放数据库资源
			connect = openDatabse()
			defer connect.close()
	*/
}

func sum(n1 int, n2 int) int {
	// 当函数执行到 defer 时，暂时不执行，会将 defer 后面的语句压入独立的栈(defer 栈)
	// 当函数执行完毕后，再从 defer 栈，按照先入后出的方式出栈
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	// defer 将语句放入到栈时，也会将相关的值拷贝同时入栈
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}
