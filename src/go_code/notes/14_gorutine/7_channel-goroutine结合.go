package main

import (
	"fmt"
	"time"
)

/*
	1. 开启 writeData 协程，向管道 intChan 写入50个整数
	2. 开启 readData 协程，从管道 intChan 读取 writeData 写入的数据
	3. writeData 和 readData 操作的是同一个管道
	4. 主线程需要等待 writeData 和 readData 协程都完成工作才能退出
*/

func writeData(intChan chan int) {
	for i := 0; i <= 10; i++ {
		intChan <- i // 放入数据
		fmt.Println("writeData", i)
		time.Sleep(time.Second * 1)
	}
	close(intChan) // 关闭管道
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	// readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	// 创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
