package main

import "fmt"

func main() {
	// slice_usage()
	slice_usage2()
}

func notes() {
	/*
		***** 切片的基本介绍 *****
		1. 切盼是数组的一个引用，因此切片是引用类型，在进行传递时，遵循引用传递的机制
		2. 切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度 len(slice) 都一样
		3. 切片的长度是可以变化的，因此切片是一个可以动态变化数组
		4. slice 从底层来说，其实就是一个 struct 结构体
	*/
}

func slice_usage() {
	// 演示切片的基本使用
	var array [5]int = [...]int{71, 23, 41, 56, 30}
	/*
		1. slice 就是切片名
		2. array[1:3] 表示 slice 引用到 array 这个数组的第下标为1的元素到下表为3的元素，不包含3
	*/
	slice := array[1:3]
	fmt.Printf("slice 的元素是%v\n", slice)
	fmt.Printf("slice 的元素个数是%v\n", len(slice))
	fmt.Printf("slice 的容量%v\n", cap(slice))

	fmt.Println(array)
	slice[0] = 100
	fmt.Println(array)
}

func slice_usage2() {
	var intArr [5]int = [...]int{11, 22, 33, 44, 55}
	// intArr[1:3] 表示 slice 引用到 intArr 这个数组，引用到 intArr 数组的起始下标为1，最后下标为3（但不包含3）
	slice := intArr[1:3]
	fmt.Println("intArr,", intArr)
	fmt.Println("slice 元素是:", slice)
	fmt.Println("slice 的元素个数", len(slice))
	fmt.Println("slice 的容量", cap(slice))
}
