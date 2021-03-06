package main

import (
	"fmt"
	"unsafe"
)

func main() {
	my_bool()
}

func my_bool() {
	/*
		1、布尔类型也叫 bool 类型，bool 类型数据只允许取值 true 和 false
		2、bool 类型只占 1个字节
		3、boolean 类型适合于逻辑运算，一般用于程序流程控制。
			if 条件控制语句
			for 条件控制语句
	*/
	var b bool = false
	fmt.Println("b=", b)
	fmt.Println("布尔类型占用的空间=", unsafe.Sizeof(b))
}
