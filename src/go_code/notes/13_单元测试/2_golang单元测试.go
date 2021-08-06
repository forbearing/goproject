package main

// 一个被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {

}

/*
golang 中带有一个轻量级的测试框架 testing 和自带的 go test 命令来实现单元测试和性能测试，
testing 框架和其它语言的测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基于
该框架写相应的压力测试用例。通过单元测试，可以解决如下问题：
	1. 确保每个函数是可运行的，并且运行结果是正确的
	2. 确保写出来的代码性能是好的
	3. 单元测试能及时的发现程序设计实现的逻辑错误，使问题及早暴露，便于问题的定位解决，而性能
	   测试的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保持稳定。
*/
