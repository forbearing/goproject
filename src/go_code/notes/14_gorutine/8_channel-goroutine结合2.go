package main

import (
	"fmt"
	"time"
)

// 向 intChan 放入数据
func putNum(intChan chan int, n int) {
	for i := 0; i <= n; i++ {
		intChan <- i
	}
	close(intChan) // 关闭 intChan，方便 read 判断读取完所有元素后退出
}

// 从 intChan 取出数据并判断是否为素数，如果是就放入到 primeChain
func primeNum(intChan chan int, primechan chan int, exitChan chan bool) {
	for {
		// time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok { // intChan 取不到
			break
		}
		flag := true // 假设不是素数
		// 判断 num 是否是素数
		for i := 2; i < num; i++ {
			if num%2 == 0 { // 说明不是素数
				flag = false
				break
			}
		}
		if flag {
			primechan <- num // 将这个数放入到 primechan
		}
	}
	fmt.Println("有一个 primeNum 协程因为取不到数据，退出")
	// 这里还不能关闭 primeChan
	// 向 exitChan 写入 true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000000)
	primeChan := make(chan int, 1000000) // 放入结果
	exitChan := make(chan bool, 4)       // 标识退出

	// 开启一个协程，向 intChan 放入 1-8000 个数
	start := time.Now().Unix()
	go putNum(intChan, 800000)

	// 开启四个协程，从 intChan 取出数据并判断是否为素数，如果是就放入到 primeChain
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 当我们从 exitChain 取出了4个结果，就可以放心的关闭 primeNum 了
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时:", end-start)
		close(primeChan)
	}()

	// 遍历我们的 primeNum，把结果取出来
	for {
		_, ok := <-primeChan
		// res,ok := <-primeChain
		if !ok {
			break
		}
		// // 将结果输出
		// fmt.Printf("%v ", res)
	}
	fmt.Printf("\nmain 线程退出\n")
}
