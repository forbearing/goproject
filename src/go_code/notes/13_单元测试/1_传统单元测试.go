package main

import "fmt"

// 一个被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	// 传统的测试方法，就是在 main 函数中使用查看结果是否正确
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper 错误，返回值=%v 期望值=%v\n", res, 55)
	} else {
		fmt.Printf("addUpper 正确，返回值=%v 期望值=%v\n", res, 55)
	}
}

/*
传统单元测试的缺点
	1. 不方便，我们需要在 main 函数中去调用，这样就需要去修改 main 函数，如果现在项目正在运行，
	   就可能要去停止项目
	2. 不利于管理，因为当我们测试多个函数或者多个模块时，都需要写在 main 函数，不利于我们管理和清晰我们的思路
	3. 引出单元测试：testing 框架，可以很好的解决问题。
*/
