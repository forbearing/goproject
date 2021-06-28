package main

import "fmt"

func main() {
	// example1()
	slice_copy()
}

func notes() {
	/*
		切片的遍历两种方式：for 循环、for-range
		***** 切片注意事项和细节说明 *****
		1. 切片初始化时 var slice = arr[startIndex:endIndex]
			从数组下标为 startIndex，取到下标为 endIndex 的元素，不包含下标
		2. 切片初始化时，仍然不能越界，范围在 0-len(arr) 之间，但是可以动态增长
		3. 简写
			var slice = arr[0:end] ==> var slice = arr [:end]
			var slice = arr[start:len(arr)] ===> var slice = arr[start:]
			var slice = arr[0:len(arr)] ==> var slice = arr[:]
		4. cap 是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
		5. 切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者 make 一个空间供切片来使用
		6. 切片可以继续切片
		7. 使用 append 内置函数，可以对切片进行动态追加
			切片 append 操作的底层原理分析
			1. 切片 append 操作的本质是对数组进行扩容
			2. go 底层会创建一个新的数组 newArr（扩容后的大小）
			3. 将 slice 原来包含的元素拷贝到新的数组 newArr
			4. slice 重新引用到新的数组
			5. 注意 newArr 是在底层维护，对程序员不可见
		8. 切片的拷贝操作
			切片使用 copy 内置函数完成拷贝
			copy(slice1, slice2), slice1 和 slice2 都是i切片类型

	*/
}

func example1() {
	// 方式一：常规 for 循环遍历
	var arr = [...]int{11, 22, 33, 44, 55}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]: %v\n", i, slice[i])
	}

	// 方式二：使用 for-range 遍历
	fmt.Println()
	for index, value := range slice {
		fmt.Printf("index=%v value=%v\n", index, value)
	}

	// 切片可以继续切片
	fmt.Println()
	slice2 := slice[1:2]
	slice2[0] = 100
	fmt.Println("slice2=", slice2)
	fmt.Println("slice", slice)
	fmt.Println("arr", arr)
}

func slice_append() {
	// 使用 append 内置函数，可以对切片进行动态追加
	var slice1 []int = []int{100, 200, 300}
	fmt.Println("slice1 =", slice1)

	// slice2 := append(slice1, 400, 500)
	// fmt.Println("slice2 =", slice2)

	slice1 = append(slice1, 400, 500, 600)
	fmt.Println("slice1 =", slice1)

	slice1 = append(slice1, slice1...) // 切片追加到自己
	// slice1 = append(slice1, slice2...)
	fmt.Println("slice1 =", slice1)

}

func test(slice []int) {
	slice[0] = 100
}

func slice_copy() {
	var slice1 []int = []int{11, 22, 33, 44, 55}
	slice2 := slice1
	var slice3 = make([]int, 10)
	copy(slice3, slice1)
	fmt.Println("slice1 =", slice1)
	fmt.Println("slice2 =", slice2)
	fmt.Println("slice3 =", slice2)
	fmt.Println("slice1 addr =", &slice1[0])
	fmt.Println("slice2 addr =", &slice2[0])
	fmt.Println("slice3 addr =", &slice3[0])

	fmt.Println()
	var slice4 = make([]int, 1)
	copy(slice4, slice1) // slice4 比 slice1 长度大，但不会报错，只会 copy 一个元素而已，代码没有错
	fmt.Println("slice1 =", slice1)
	fmt.Println("slice4 =", slice4)

	// 切片是引用类型，所以在传递时，遵循引用传递机制。
	fmt.Println()
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println("slice before =", slice)
	test(slice)
	fmt.Println("slice after =", slice)
}
