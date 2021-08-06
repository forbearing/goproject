package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex // 声明一个全局的互斥锁，lock 是一个全局的互斥锁
)

// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map 中。
// 2. 启动协程多个，统计的结果放入到 map 中。
// 3. map 应该做成一个全局的。

// test 函数就是计算 n!，将这个结果放入到 map 中
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func main() {
	// 这里开启多个协程来完成这个任务
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	// 即使程序可以10秒内运行完，不应该出现资源竞争的情况，但是主线程并不知道，因此底层可能仍然出现资源竞争，因此加入互斥锁即可解决资源竞争问题。
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}

func notes() {
	/*
		===== 不同 goroutine 之间如何通信
		1.全局变量加锁同步
		2.channel
	*/
}
