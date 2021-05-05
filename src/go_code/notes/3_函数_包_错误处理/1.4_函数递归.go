package main

import "fmt"

func main() {
	// test2(4)
	//res := fbn(10)
	//fmt.Println(res)

	//num := 10
	//res = changeMe(&num)
	//fmt.Println(res)
	//fmt.Println(num)

	num1 := 10
	num2 := 20
	res := myFunc(getSum, num1, num2)
	fmt.Println(res)
}

func notes() {
	/*
		一个函数在函数体内又调用了本身，称为递归调用。

		*** 函数递归需要遵守的重要原则 ***
		1. 执行一个函数时，就创建一个新的受保护的独立空间（新函数栈）
		2. 函数的局部变量是独立的，不会相互影响。
		3. 递归必须向退出递归的条件逼近，否则就无限递归了
		4. 当一个函数执行完毕，或者遇到 return，就会返回，遵守谁调用，就将结果返回给谁
		   同时当函数执行完毕或者返回时，该函数本身也会被销毁。
	*/
}

func test1(n int) {
	if n > 2 {
		n--
		test1(n)
	}
	fmt.Println("n=", n)
}
func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n=", n)
	}
}

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func changeMe(num *int) int {
	*num++
	return *num
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func myFunc(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}
