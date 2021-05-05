package main

import "fmt"

func main() {
	my_func1()
}

func notes() {
	/*
		***** 介绍 *****
		1. Go 语言的 goto 语句可以无条件的跳转到程序中指定的行
		2. goto 语句通常与条件语句配合使用。可以来实现条件转移，跳出循环体等功能。
		3. 在 Go 程序设计中一般不主张使用 goto 语句，以免造成程序流程的混乱，使
		   理解和调试程序都产生困难。
	*/
}
func my_func1() {
	// goto 的使用
	fmt.Println("ok1")
	fmt.Println("ok2")
	goto label4
	fmt.Println("ok3")
label4:
	fmt.Println("ok4")
}
