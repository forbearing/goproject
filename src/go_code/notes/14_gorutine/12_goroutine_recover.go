package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World")
		time.Sleep(time.Second)
	}
}

func test() {
	// 使用 defer + recover
	defer func() {
		// 捕获 test 函数抛出的 panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误")
		}
	}()
	// 定义一个map
	var myMap map[int]string
	myMap[0] = "golang" // 没有 make，错误
}

func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("main() Ok")
	}
}
