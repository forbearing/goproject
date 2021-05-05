package main

import "fmt"

func main() {
	// my_func1()
	// res := cal(1, 10, '+')
	// fmt.Println(res)
	//myType()

	// sum := sum1(1, 2, 3, 4, 5)
	// fmt.Println(sum)

	num1 := 3
	num2 := 4
	fmt.Println(num1, num2)
	swap(&num1, &num2)
	fmt.Println(num1, num2)

}

func notes() {
	/*
		***** 概述 *****
		1. 为完成某一功能的程序指令(语句)的集合，称为函数。
		2. 在 Go 中，函数分为：自定义函数，系统函数

		***** 基本语法 *****
		func 函数名 (形参列表) (返回值类型列表) {
			执行语句 ...
			return 返回值列表
		}
		1. 形参列表：表示函数的输入
		2. 函数中的语句：表示为了实现某一功能代码块
		3. 函数可以有返回值，也可以没有

		***** 函数使用细节和注意事项 ******
		1. 函数的形参列表可以是多个，返回值列表也可以是多个
		2. 形参列表和返回值列表的数据类型可以是值类型和引用类型
		3. 函数的命名遵循标识符命名规范，首字母不能是数字，首字母大写该函数可以被本包文件和其它
		   包文件使用，类似 public，首字母小写，只能被本包文件使用，其它包文件不能使用，类似 private
		4. 函数中的变量是局部的，函数外不生效
		5. 基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值。
		6. 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址 &，函数内以指针的方式操作变量。
		   从效果上看类似引用。
		7. Go 函数不支持重载
		8. 在 Go 中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，
		   通过该变量可以对函数调用
		9. 函数既然是一种数据类型，因此在 Go 中，函数可以作为形参，并且调用
		10. 为了简化数据类型定义，Go 支持自定义数据类型
			// 给函数取个别名
			type myFunType func(int, int) int
			func myFunc(funvar func(int, int) int, num1 int, num2 int) int {
				return funvar(num1, num2)
			}
			func myFunc2(funvar myFunType, num1 int, num2 int) int {
				return funvar(num1, num2)
			}
		11. 支持对函数返回值命名
			func cal1(n1 int, n2 int) (sum int, sub int) {
				sum = n1 + n2
				sub = n1 - n2
				return
			}
		12: 使用 _ 标识符，忽略返回值
		13. Go 支持可变参数
			1. 支持 0到多个参数
				func sum(args... int) sum int {}
			2. 支持1到多个参数
				func sum(n1 int, args... int) sum int {}
			3. args 是 slice，通过 args[index] 可以访问到各个值
			4. 如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后。

	*/
}

func my_func1() {
	// 传统方法，不使用函数
	var num1 float64
	var num2 float64
	var operator string
	var res float64
	fmt.Printf("操作符：[+|-|*|/]: ")
	fmt.Scanf("%s", &operator)
	fmt.Printf("第一个数: ")
	fmt.Scanf("%f", &num1)
	fmt.Printf("第二个数: ")
	fmt.Scanf("%f", &num2)

	switch operator {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2

	default:
		fmt.Println("操作符错误")
	}
	fmt.Println("结果是:", res)
}
func cal(num1 float64, num2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	case '/':
		res = num1 / num2
	default:
		fmt.Println("操作符错误...")
	}
	return res
}
func myType() {
	// 使用自定义类型
	// 给 int 取了个别名，在 go 中，myInt 和 int 虽然都是 int 类型，
	// 但是 go 认为 myInt 和 int 是不同的两个类型
	type myInt int
	var num1 myInt
	var num2 = 40
	num1 = 40
	num2 = int(num1) // 需要类型转换，go 认为 int 和 myInt 是两种类型。

	fmt.Println(num1)
	fmt.Println(num2)
}

// 支持对函数返回值命名
func cal1(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 函数形参的可变参数
// 编写一个函数，求出1到多个数的和
func sum1(n1 int, args ...int) int {
	sum := n1
	// 遍历 args
	for i := 0; i < len(args); i++ {
		sum += args[i] // args[0] 表示取出 args 切片的第一个元素值.
	}
	return sum
}

// 交换两个变量的值
func swap(num1 *int, num2 *int) {
	t := *num1
	*num1 = *num2
	*num2 = t
}
