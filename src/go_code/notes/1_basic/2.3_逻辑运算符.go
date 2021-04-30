package main

import "fmt"

func main() {
	my_func()
}

func my_func() {
	/*
		1. 介绍
			1. 用于连接多个条件（一般来说关系表达式），最终结果也是一个 bool 值。
			2. && 逻辑与
			3. || 逻辑或
			4. ! 逻辑非
		2. 注意事项和细节说明
			1. && 也叫短路与，如果一个条件为 false，则第二个条件不会判断，最终结果为 false
			2. || 也叫短路或，如果第一个条件为 true，则第二个条件不会判断，最终结果为 true
	*/
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("30~50")
	} else if age < 30 {
		fmt.Print("< 30")
	} else {
		fmt.Print("> 50")
	}

	if age < 30 || age < 40 {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	if !(age > 0) {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}
