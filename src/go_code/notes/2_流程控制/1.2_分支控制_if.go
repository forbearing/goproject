package main

import (
	"fmt"
	"math"
)

func main() {
	//my_func()
	// my_func2()
	// my_func3()
	my_func4()
}

func my_func() {
	/*
		分支控制 if-else
		让程序有选择的执行，分支控制有三种
		单分支、
		双分支、
		多分枝、
	*/

	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你的年龄大于18岁，你要对你的行为负责")
	} else {
		fmt.Println("你的年龄小于18岁")
	}

	// Go 的 if 还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域
	// 只能在该条件逻辑块内，其它地方就不起作用。
	if age := 20; age > 18 {
		fmt.Println("年龄大于18岁")
	}

}

func my_func2() {
	var n1 float32 = 11.0
	var n2 float32 = 21.1
	if n1 > 10 && n2 < 30 {
		fmt.Println("ok")
	}

	var n3 int32 = 3
	var n4 int32 = 5
	if 15%n3 == 0 && 15%n4 == 0 {
		fmt.Println("ok")
	}

	// 判断一个年份是否是闰年：
	//	1.年份能被4整除但不能被100整除
	//	2.能被400整除
	var year int = 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("是闰年")
	}
}

// 多分支
func my_func3() {
	var score int
	fmt.Println("请输入成绩:")
	// fmt.Scanln(&score)
	fmt.Scanf("%d", &score)
	if score == 100 {
		fmt.Println("奖励宝马一台")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台 iPhone")
	} else if score > 60 && score <= 80 {
		fmt.Println("奖励一台 iPad")
	} else {
		fmt.Println("什么奖励也没有")
	}

	var a float64 = 3.0
	var b float64 = 100.0
	var c float64 = 6.0

	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v x2=%v", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v x2=%v", x1, x1)
	} else {
		fmt.Println("无解...")
	}
}

// 嵌套分支
func my_func4() {
	/*
		嵌套分支不宜过多，最多我们建议控制在3层。
	*/

	var second float64
	fmt.Println("请输入秒数")
	fmt.Scanln(&second)
	if second <= 8 {
		// 进入决赛
		var gender string
		fmt.Println("请输入你的性别")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入决赛男子组")
		} else {
			fmt.Println("进入决赛的女子组")
		}
	} else {
		fmt.Println("out ...")
	}
}
