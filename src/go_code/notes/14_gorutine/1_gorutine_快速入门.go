package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() // 开启了一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

/*
需求：统计 1-100000000 的数字中，哪些是素数
分析思路：
	1.传统方法，使用一个循环，循环的各个数是不是素数
	2.使用并发或并行的方式，将统计素数的任务分配给多个 goroutine 去完成，这时就会使用到 goroutine

goroutine 基本介绍
	1.多线程程序在单核上运行，就是并发
	2.多线程陈旭在多核上运行，就是并行

Go 协程和 Go 主线程
	1.Go主线程（有程序员直接称线程/也可以理解成进程）：一个 Go 线程上，可以起多个协程，
	  可以这样理解，协程就是轻量级的线程。
	2.Go 协程的特点：
		有独立的的栈空间
		共享程序堆空间
		调度由用户控制
		协程是轻量级的线程

goroutine 小结：
	1.主线程是一个物理线程，直接作用于cpu上的。是重量级的，非常消耗资源。
	2.协程是从主线程开启的，是轻量级的，是逻辑态的，对资源消耗相对小
	3.Golang 的协程机制是重要的特点，可以轻松的开启上万个协程。其它编程语言的并发机制一般是基于线程的，
	  开启过多，资源耗费大，这里就突显了 Golang 在并发上的优势了。
*/
