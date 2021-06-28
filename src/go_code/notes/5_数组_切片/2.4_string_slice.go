package main

import "fmt"

func main() {
	example()
}

func notes() {
	/*
		***** string & slice *****
		1. string 底层是一个 byte 数组，因此 string 也可以进行切片处理
		2. string 是不可变的，不能通过 string[0] = 'z 方式来修改字符串
		3. 如果需要修改字符串，可以先将 string -> []byte 或者 []rune -> 修改重写转成 string
	*/
}

func example() {
	str := "com.hybfkuf"
	slice := str[4:]
	fmt.Println("slice =", slice)

	// 修改 string
	// 转成 []byte 后，可以处理英文和数字，但不能处理中文，解决办法就是将 string 转成 []rune 即可
	arr := []byte(str)
	arr[0] = 'z'
	str = string(arr)
	fmt.Println("str =", str)

	arr1 := []rune(str)
	arr1[0] = '中'
	str = string(arr1)
	fmt.Println("str =", str)
}
