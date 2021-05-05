package main

import "fmt"

func main() {
	my_func1()
	my_func2()
}
func notes() {
	/*
		1. 将一个循环放在另外一个循环体内，就形成了嵌套循环。在外面的 for 称为外层循环，在里面的 for 循环
		   称为内层循环。建议一般就两层，最多不要超过3层。
		2. 实质上，嵌套循环就是把内层循环当成外层循环的循环体。当只有内层循环的循环条件为 false 时，才会
		   完全跳出内层循环，才可结束外层的循环
		3. 设外层循环次数为 m 次，内层为 n 次，则内层循环体实际上需要执行 m*n=mn 次。
	*/
}

var totalLevel int = 10

func my_func1() {
	// 打印金字塔
	for i := 1; i <= totalLevel; i++ {
		// 打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func my_func2() {
	// 打印九九乘法表
	var num int = 4
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, j*i)
		}
		fmt.Printf("\n")
	}
}
