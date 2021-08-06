package main

import "fmt"

func main() {
	// 1. 创建一个可以存放3个 int 类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 2. 看看 intChan 是什么
	fmt.Printf("intChan 的值 %v\nintchan 本身的地址%v\n", intChan, &intChan)

	// 3. 向管道写入数据
	intChan <- 10
	num1 := 20
	intChan <- num1
	intChan <- 30
	// intChan <- 30 // 当我们给管道写入数据时，不能超过其容量

	// 4. 看看管道的长度和cap（容量）
	fmt.Printf("channnel len =%v, cap=%v\n", len(intChan), cap(intChan))

	// 5.从管道读取数据
	num2 := <-intChan
	num3 := <-intChan
	num4 := <-intChan
	fmt.Println(num2, num3, num4)
	// num5 := <-intChan
	// fmt.Println(num5) //在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock

}

/*
1. 为什么需要 channel
	1. 主线程在等待所有 goroutine 全部完成的时候很难确定
	2. 如果主线程休眠时间长，会加长等待时间，如果等待时间短了，可能还有 goroutine 处于工作状态，这时也会随主线程的退出而销毁。
	3. 通过全局变量加锁同步来实现通讯，也并不利多个协程对全局变量的读写操作
2. channel 的介绍
	1. channel 本质就是一种数据结构-队列
	2. 数据是先进先出
	3. 线程安全，多 goroutine 访问时，不需要加载，就是说 channel 本身就是线程安全的
	4. channel 是有类型的，一个 string 的 channel 只能存放 string 类型数据

1. channel 基本使用
	定义/声明 channel
		var 变量名 chan 数据类型
		var intChan chan int (intChan 用于存放 int 数据)
		var mapChan chan map[int]string (mapChan 用于存放 map[int]string 类型)
		var perChan chan Person
		var perChan2 chan *Person
	说明
		1. channel 是引用类型
		2. channel 必须初始化才能写入数据，即 make 之后才能使用
		3. 管道是有类型的，intChan 只能写入整数 int

2. channel 使用注意事项
	1. channel 只能存放指定的数据类型
	2. channel 的数据放满后，就不能再放了。
	3. 如果从 channel 取出数据后，可以继续放入
	4. 在没有使用协程的情况下，如果 channel 数据取完了，再取，就会报 dead lock
*/
