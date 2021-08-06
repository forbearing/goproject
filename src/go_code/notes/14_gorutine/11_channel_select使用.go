package main

import (
	"fmt"
	"time"
)

func main() {
	// 使用 selector 可以解决从管道读取数据的阻塞问题
	// 1. 定义一个管道 10 个数据 int

	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// 2. 定义一个管道 5 个数据 string
	stringchan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringchan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统方法在遍历管道时，如果没有关闭管道，会导致阻塞而 deadlock
	// 问题：在实际开发中，我们不好确定什么时候关闭管道，可以使用 selector 方式来解决
label:
	for {
		select {
		case v := <-intChan: // 注意：如果 intChan 一直没有关闭，不会因为阻塞导致 deadlock，会自动到下一个 case 匹配
			fmt.Printf("从 intChan 读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringchan:
			fmt.Printf("从 stringChan 读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("读取不到了, 不玩了")
			break label
		}
	}
}
