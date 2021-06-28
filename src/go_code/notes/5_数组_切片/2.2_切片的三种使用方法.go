package main

import "fmt"

func main() {
	slice_usage()
}

func notes() {
	/*
		***** 切片的使用 *****
		1. 方式一：定义一个切片，然后让切片去引用一个已经创建好的数组
			var arr [5]int = [...]int {1,2,3,4,5}
			var slice = arr[1:3]
		2. 方式二：通过 make 来创建切片
			var slice []int = make([]int, len, [cap])
			var slice []int = make([]int, 4, 10)
			cap 指定切片容量，可选，如果指定了 cap，则要求 cap >= len
			* 内建函数 make 分配并初始化一个类型为切片、映射、或通道的对象。其第一个实参为类型，而非值。
			  make 的返回类型与其参数相同，而非指向它的指针。
		3. 方式三：定义一个切片，直接就指定具体数组，使用原理类似 make 的方式
			var slice []int = []int {1,3,5}
			fmt.Println(slice)

		***** 方式一和方式二的区别 *****
		- 方式一是直接引用数组，这个数组是事先存在的，程序员是可见的。
		- 方式二是通过 make 来创建切片，make 也会创建一个数组，是由切片在底层进行维护，程序员是看不见的。

	*/
}

func slice_usage() {
	// 方式二：
	// var slice [] float64		数组可以这么用， slice 不可以
	// 对于切片必须 make 之后才能使用。
	// 通过 make 方式创建切片可以指定切片大小和容量，如果没有给切片的各个元素赋值，就是用默认值。
	// 通过 make 方式创建的切片对应的数组是由 make 底层维护，对外不可见。只能通过 slice 访问
	var slice []float64 = make([]float64, 5, 10)
	fmt.Println(slice)
	fmt.Printf("slice 的大小: %v\n", len(slice))
	fmt.Printf("slice 的容量: %v\n", cap(slice))

	// 方式三：
	// 定义一个切片，直接就指定具体数组，使用原理类似 make 的方式
	fmt.Println()
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice =", strSlice)
	fmt.Println("strSlice size =", len(strSlice))
	fmt.Println("strSlice cap =", cap(strSlice))
}
