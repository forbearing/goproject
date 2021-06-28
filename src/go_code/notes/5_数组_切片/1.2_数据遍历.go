package main

import "fmt"

func main() {
	for_range_example()

}

func notes() {
	/*
		***** 数组的遍历 *****
		方法一: 常规遍历
		方法二: for-range 结构遍历
			这是 Go 语言一种独有的结构，可以用来遍历访问数组的元素。
			for index, value := range array {...}
			1. 第一个返回值 index 是数组的下表
			2. 第二个 value 是在该下标位置的值
			3. 他们都是仅在 for 循环内部可见的局部变量
			4. 遍历数组元素的时候，如果不想使用下标 index，可以直接下标 index标为下划线 _
			5. index 和 value 的名称不是固定的，程序员可以自行指定，一般为 index 和 value

	*/
}

func for_range_example() {
	heroes := [...]string{"宋江", "吴用", "卢俊义"}
	for index, value := range heroes {
		fmt.Printf("index=%v value=%v\n", index, value)
	}
	for _, value := range heroes {
		fmt.Println(value)
	}
}
