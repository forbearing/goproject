package main

import (
	"fmt"
	"os"
)

// 编写一段代码，获取命令行各个参数

func main() {
	fmt.Println("所有的命令行参数有", len(os.Args))
	// 遍历 os.Args 切片，就可以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v\n", i, v)
	}
}

/*
os.Args 是一个 string 的切片，用来存储所有的命令行参数
*/
