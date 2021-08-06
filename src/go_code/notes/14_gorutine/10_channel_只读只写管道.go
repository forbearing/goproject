package main

import (
	"fmt"
	"time"
)

func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond * 10)
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		time.Sleep(time.Microsecond * 10)
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}

func main() {
	var ch chan int
	ch = make(chan int, 10)
	exitchan := make(chan struct{}, 2)
	go send(ch, exitchan)
	go recv(ch, exitchan)

	total := 0
	for _ = range exitchan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束 ...")
}
