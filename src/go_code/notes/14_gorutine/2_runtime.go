package main

import (
	"fmt"
	"runtime"
)

// go 1.8 之后，默认让程序运行在多核之上，可以不用设置
// go 1.8 前，需要设置，可以更高效的利用 cpu

/*
计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到 map 中。最后显示出来，使用 goroutine 完成
分析：
	1.使用 goroutine 来完成，效率高，但是会出现并发/并行安全问题。
	2.goroutine 如何通信
*/

// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map 中。
// 2. 启动协程多个，统计的结果放入到 map 中。
// 3. map 应该做成一个全局的。

var (
	myMap = make(map[int]int, 10)
)

// test 函数就是计算 n!，将这个结果放入到 map 中
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}
func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("你的 cpu 个数:", cpuNum)

	// // 可以自己设置使用多少个 cpu
	// runtime.GOMAXPROCS(cpuNum - 1)

	// 这里开启多个协程来完成这个任务
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
