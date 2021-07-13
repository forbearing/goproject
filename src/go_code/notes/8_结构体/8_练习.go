package main

import "fmt"

type MethodUtils struct {
}

// 1、编写结构体 MethodUtils，编写一个方法，方法不需要参数，在方法中打印一个 10*8 的矩形，在 main 方法中调用该方法
func (mu MethodUtils) Print() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 2、编写一个方法，提供 m 和 n 两个参数，方法中打印一个 m*n 的矩形
func (mu MethodUtils) Print2(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//3、编写一个方法算该矩形的面积（可以接收长 len、宽 width），将其作为方法返回值。在 main 方法中调用该方法，接收返回的面积值并打印
func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

//3、编写方法：判断一个数是奇数还是偶数
func (mu MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶数...")
	} else {
		fmt.Println(num, "是奇数...")
	}
}

// 4、根据行、列、字符打印对应行数和列数的字符，比如：行：3，列：2，字符 *，则打印相应的效果
func (mu MethodUtils) Print3(m int, n int, key string) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

// 5、定义一个计算器结构体，实现加减乘除
type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		if calcuator.Num2 == 0 {
			res = 0
			return res
		}
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("运算符输入有问题 ...")
	}
	return res
}

func main() {

	var mu MethodUtils
	mu.Print()
	fmt.Println()
	mu.Print2(8, 10)
	areaRes := mu.area(10, 8)
	fmt.Println("面积为:", areaRes)
	mu.JudgeNum(99)
	mu.Print3(6, 3, "Linux")

	res := Calcuator{
		Num1: 10,
		Num2: 0,
	}
	sum := res.getRes('/')
	fmt.Println("sum =", sum)
}

/*
1、编写结构体 MethodUtils，编写一个方法，方法不需要参数，在方法中打印一个 10*8 的矩形，在 main 方法中调用该方法
2、编写一个方法，提供 m 和 n 两个参数，方法中打印一个 m*n 的矩形
3、编写一个方法算该矩形的面积（可以接收长 len、宽 width），将其作为方法返回值。在 main 方法中调用该方法，接收返回的面积值并打印
*/

/*
3、编写方法：判断一个数是奇数还是偶数
4、根据行、列、字符打印对应行数和列数的字符，比如：行：3，列：2，字符 *，则打印相应的效果
5、定义一个计算器结构体，实现加减乘除
	实现方法一：分别用4个方法
	实现方法二：用一个方法
*/
