package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// my_func1()
	example()
}

func notes() {
	/*
		break 语句用于终止某个语句块的执行，用于中断当前 for 循环或跳出 switch 语句。

		***** 注意事项和细节说明 *****
		1. break 语句出现在多层嵌套的语句块中时，可以通过标签指明要终止的是哪一层语句块
		2. break 默认跳出最近的 for 循环, break 后面可以指定标签，跳出标签对应的 for 循环。
		3. 标签的基本使用


	*/
}

func my_func1() {
	// 随机生成 1-100 的一个数，直到生成 99 这个数，看一共使用了多少次。
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			break
		}
	}
	fmt.Printf("生成99，一共用了 %d 次\n", count)

	// break 跳出多从循环
label1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label1
			}
			fmt.Println("j=", j)
		}
	}
}
func example() {
	var name string
	var pass string
	var changeCount int = 3
	count := changeCount
	for i := 0; i < count; i++ {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&name)
		fmt.Println("请输入你的密码")
		fmt.Scanln(&pass)
		if name == "hybfkuf" && pass == "888" {
			fmt.Println("登录成功")
			break
		} else {
			changeCount--
			if changeCount == 0 {
				fmt.Println("你已经没机会了，登录失败")
			} else {
				fmt.Printf("你还有 %d 次机会\n", changeCount)
			}
		}
	}
}
