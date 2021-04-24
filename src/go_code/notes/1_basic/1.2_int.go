package main

import (
	"fmt"
	"unsafe"
)

func main() {
	integer()
}

// 有符号整数使用
func integer() {
	// int8 int16 int32 int64 依此类推
	var i1 int8 = -128
	fmt.Println("i1=", i1)
	//var i2 int8 = -129 // int8 最小为 -128, 最大为 127
	//fmt.Println("i2=", i2)

	// uint8 uint16 uint32 uint64
	// uint8 的范围 0-255。

	// rune 等价于 int32，表示一个 unicode 编码
	// byte，等价于 uint8，当要存储字符时选用 byte
	var char byte = 255
	fmt.Println("char=", char)

	// 输出变量的数据类型
	var n1 int64 = 100
	fmt.Printf("n1 的数据类型: %T, n1 占用的字节 数: %d", n1, unsafe.Sizeof(n1))

	// 整型的使用细节
	// 1、Golang 的各整数类型分：有符号和无符号，int uint 的大小和系统有关。
	// 2、Golang 的整型默认声明为 int 型
	// 3、如何在程序查看某个变量的字节大小和数据类型（使用较多）
	// 4、Golang 程序中整型变量在使用时，遵循保小不保大的原则，即：在保证程序正常运行下，尽量使用占用空间小的数据类型
	var age byte = 99
	fmt.Println("age=", age)
	// 5、bit：计算机中的最小存储单位，byte：计算机中基本存储单元
}
