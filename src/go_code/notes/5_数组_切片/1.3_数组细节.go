package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// arr := [3]int{11, 22, 33}
	// pass_array_by_value(arr)
	// fmt.Println(arr)
	// pass_array_by_reference(&arr)
	// fmt.Println(arr)

	// exmaple1()
	exmaple2()
}

func notes() {
	/*
		**** 数组使用注意事项和细节 *****
		1. 数组是多个相同类型的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化的。
		2. var arr[]int 这时 arr 就是一个 slice 切片
		3. 数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用
		4. 数组创建后，如果没有赋值，有默认值
			1. 数值类型数组，默认值为0
			2. 字符串睡着，默认为 ""
			3. bool 数组，默认为 false
		5. 使用数组的步骤：
			1. 声明数组并开辟空间
			2. 给数组各个元素赋值
			3. 使用数组
		6. 数组的下标是从0开始的
		7. 数组下标必须在指定范围内使用，否则报 panic 错误，数组越界
		8. Go 的数组属于值类型，在默认情况下是值传递，因此会进行拷贝。数组间不会相互影响。
		9. 如想在其他函数中，去修改原来的数组，可以使用引用传递(指针方式)
		10. 长度是数组的一部分，在传递参数时，需要考虑数组的长度。
	*/
}

func pass_array_by_value(arr [3]int) {
	arr[0] = 88
}
func pass_array_by_reference(arr *[3]int) {
	// arr[1] = 88
	(*arr)[1] = 88
}
func exmaple1() {
	var myChars [26]byte
	myChars[0] = 'A'
	for i := 1; i < 26; i++ {
		myChars[i] = myChars[0] + byte(i)
	}

	for _, value := range myChars {
		fmt.Printf("%c ", value)
	}
	fmt.Println()
}
func exmaple2() {
	// 数组反转
	var array1 [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array1); i++ {
		array1[i] = rand.Intn(100)
	}
	fmt.Println(array1)
}
