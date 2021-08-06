package main

import "fmt"

func main() {
	// 1.管道可以声明为只读或者只写
	//   在默认情况下，管道是双向的
	var chan1 chan int   // 可读可写
	var chan2 chan<- int // 声明为只写
	var chan3 <-chan int // 声明为只读
	chan1 = make(chan int)
	chan2 = make(chan int)
	chan3 = make(chan int)
	fmt.Println(chan1)
	fmt.Println(chan2)
	fmt.Println(chan3)
}

/*
channel 使用细节和注意事项
	1.channel 可以声明为只读，或者只写性质
	2.使用 select 可以解决从管道读取数据的阻塞问题
	3.goroutine 中使用 recover，解决协程中出现 panic，导致程序奔溃问题
	  如果我们起了一个协程，但是这个协程出现了 panic，如果我们没有捕获这个 panic，
	  就会造成整个程序崩溃，这时可以在 goruntine 中使用 recover 来捕获 panic，进行处理，
	  这样即使这个协程发生问题，但是主线程仍然不受影响，可以继续运行。
*/
