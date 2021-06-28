package main

import "fmt"

func main() {
	my_map()
}

func notes() {
	/*
		*** map ***
		map 是 key-value 数据结构，又称为字段或者关联数组。类似其它编程语言的集合。在编程中经常使用到。
		***** map 的声明 *****
			- var map 变量名 map[keyType]valueType
			- golang 中的 map，的 key 可以是很多类型，比如 bool、数字、string、指针、channel，还可以是包含
			  前面几个类型的接口、结构体、数组，通常为 int、string。注意：slice、map 还有 function 不可以，因为这几个没法用 == 来判断
			- valueType 的类型和 keyType 基本一样，通常为：数字（整数、浮点数）、string、map、struct
			- 举例
				var a map[string]string
				var a map[string]int
				var a map[int]string
				var a map[string]map[string]string
				注意：声明不会分配内存的，初始化需要 make，分配内存后才能赋值和使用。

		***** map 注意事项 *****
			1. map 在使用前一定要 make
			2. map 的 key 是不能重复，如果重复了，则以最后这个 key-value 为准
			3. map 的 value 是可以相同的。
			4. map 的 key-value 是无序的
			4. make 给 map 可以不指定 size，默认 size 为0。

		***** map 的使用方式 *****
			1.
	*/
}

func my_map() {
	// map 的声明
	var a map[string]string
	fmt.Println("map a =", a)

	// 在使用 map 前，需要先 make，make 的作用就是给 map 分配数据空间。
	b := make(map[string]string, 10)
	b["name"] = "hybfkuf"
	b["age"] = "18"
	b["gender"] = "man"
	fmt.Println("map b =", b)
}
