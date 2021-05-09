package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// myTime()
	startTime := time.Now().Unix()
	example()
	endTime := time.Now().Unix()
	fmt.Printf("耗费时间为 %v 秒\n", endTime-startTime)
}

func notes() {
	/*
		time 包提供了时间的显示和测量用的函数。
	*/

}

func myTime() {
	// 1. 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v, now type=%T\n", now, now)

	// 2. 通过 now 获取年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 3. 格式化日期和时间
	// 方法一: fmt.Printf, fmt.Sprintf
	fmt.Printf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d \n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dataStr := fmt.Sprintf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d \n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dataStr)
	// 方法二: now.Format
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	// 4. 时间的常量
	/*
		const (
			Nanosecond Duration = 1
			Microsecond 		= 1000 Nanosecond
			Millisecond 		= 1000 Microsecond
			Second 				= 1000 MilliSecond
			Minute
			Hour
		)
		常量的作用：在程序中可用于获取指定时间单位的时间，比如想获得 100 毫秒: 100 * Millisecond
	*/
	// 每隔2秒输出内容
	for i := 0; ; i++ {
		fmt.Println(i)
		// time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 500)
		if i == 5 {
			break
		}
	}

	// 5. 获取当前 Unix 时间戳和 UnixNano 时间戳（作用是可以获取随机数字)
	fmt.Printf("now.Unix()=%v now.Unix()'s Type=%v\n", now.Unix(), now.Unix())
	fmt.Printf("now.UnixNano()=%v now.UnixNano()'s Type=%v\n", now.UnixNano(), now.UnixNano())
}

func example() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
}
