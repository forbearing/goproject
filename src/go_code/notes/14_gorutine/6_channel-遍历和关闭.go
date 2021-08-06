package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) // close 是一个内置函数，可以用来关闭 channel

	// 关闭管道后是不可以写数据的
	// intChan <- 300

	// 关闭管道后还可以读取数据
	n1 := <-intChan
	fmt.Println(n1)

	// 遍历管道不能使用普通的 for 循环，使用 for-range。for-range channel 只返回一个值。
	// 如果管道没有关闭，for-range 管道会出现 deadlock
	// 如果关闭了管道，则会正常的遍历数据
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	close(intChan2)
	for v := range intChan2 {
		fmt.Printf("%v\t", v)
	}

}

/*
channel 的关闭
	使用内置 close 可以关闭 channel，当 channel 关闭后，就不能再向 channel 写数据了，但是仍然可以从该 channel 读取数据
channel 的遍历
	channel 支持 for-range 的方式遍历，需要注意两个细节
	1. 在遍历时，如果 channel 没有关闭，则会出现 deadlock 的错误
	2. 在遍历时，如果 channel 已经关闭，则会正常遍历数据，遍历完后，就会退出遍历

*/
