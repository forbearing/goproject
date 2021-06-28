package main

import "fmt"

func main() {
	// example1()
	// init_array_default_value()
	// exmaple2()
	four_way_to_initialize_an_array()
}

func notes() {
	/*
		***** 数组介绍 *****
		1. 数组可以存放多个同一类型数据，数组也是一种数据类型，在 Go 中，数组是值类型。
		2. 数组的地址可以通过数组名来获取 &array
		3. 数组的第一个元素的地址就是数组的地址
		4. 数组的各个元素的地址间隔就是依据数组的类型来决定

		***** 数组的使用 ****
		4 种数组初始化的方式

	*/
}
func init_array_default_value() {
	var intArray [3]int
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	var floatArray [3]float64
	for i := 0; i < len(floatArray); i++ {
		fmt.Println(floatArray[i])
	}

	var boolArray [3]bool
	for i := 0; i < len(boolArray); i++ {
		fmt.Println(boolArray[i])
	}

	var myArray [3]int
	fmt.Printf("数组的地址: %p\n", &myArray)
	for i := 0; i < len(myArray); i++ {
		fmt.Printf("第%v个元素的地址:%p\n", i+1, &myArray[i])
	}
}
func example1() {
	// 1. 定义数组
	// var hens [6]float64;
	var hens [6]float64
	// 2. 给数组中的每个元素复制
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	fmt.Printf("总体重: %v\t平均体重:%v\n", totalWeight, totalWeight/float64(len(hens)))
}
func exmaple2() {
	// 从终端循环输入5个成绩，并保存到 float64 数组中
	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入%d个元素的值\n", i+1)
		fmt.Scanln(&score[i])
	}
	for i := 0; i < len(score); i++ {
		fmt.Println(score[i])
	}
}
func four_way_to_initialize_an_array() {
	// 四种初始化数组的方式
	var array1 [3]int = [3]int{1, 2, 3}
	fmt.Println("第一种方式", array1)
	var array2 = [3]int{1, 2, 3}
	fmt.Println("第二种方式", array2)
	var array3 = [...]int{1, 2, 3}
	fmt.Println("第三种方式", array3)
	var array4 = [...]string{1: "hybfkuf", 0: "jonas", 2: "hyb"}
	fmt.Println("第四种方式", array4)
}
