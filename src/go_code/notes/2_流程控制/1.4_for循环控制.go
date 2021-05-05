package main

import "fmt"

func main() {
	// my_func1()
	example()
}

func notes() {
	/*
		for 循环变量初始化; 循环条件; 循环变量迭代{
			循环操作语句
		}

		***** for 循环四个要素 *****
		1. 循环变量初始化
		2. 循环条件
		3. 循环操作（语句），也叫循环体
		4. 循环变量迭代

		***** for 循环执行的顺序分析 *****
		1. 执行循环变量初始化，比如 i := 1
		2. 执行循环条件判断，比如: i <= 10
		3. 如果循环条件为真，就执行循环体，比如 fmt.Println("hello golang")
		4. 执行循环变量迭代, 比如 i++
		5. 反复执行 2，3，4 步骤，直到循环条件为 false，就退出 for 循环。

		***** 注意事项和细节说明 *****
		1. 循环条件是返回一个布尔值的表达式
		2. for 循环的第二种使用方式
		   for 循环判断条件 {
			   循环执行语句
		   }
		   将变量初始化和变量迭代写到其它位置
		3. for 循环的第三种使用方式
		   for {
			   循环执行语句
		   }
		   等价于 for ;; {} 是一个无限循环，通常需要配合 break 语句使用。
		4. Golang 提供 for-range 的方式，可以方便遍历字符串和数组
		5. 如果我们的字符串含有中文，那么传统的遍历字符串方式，就会错误，会出现乱码。原因是传统的
		   对字符串的遍历是按照字节来遍历，而一个汉字在 utf8 编码时对应3个字节。
		   如何解决，需要将 str 转成 []rune 切片。
		6. for-ragne 遍历，是按照字符方式遍历，因此如果字符有中文，也是 ok 的。
	*/
}

func my_func1() {
	// for 循环的基本使用
	for i := 0; i < 3; i++ {
		fmt.Println("hello golang", i)
	}

	// for 循环的第二种使用方式
	i := 0
	for i < 3 {
		fmt.Println("hello linux", i)
		i++
	}

	// for 循环的第三种使用方法
	j := 0
	for {
		fmt.Println("hello hyb", j)
		j++
		if j == 3 {
			break
		}
	}

	// for 遍历字符串
	// 字符串遍历方式1：传统方式
	var mystr string = "my name is hybfkuf"
	for i := 0; i < len(mystr); i++ {
		fmt.Printf("%c", mystr[i])
	}
	fmt.Printf("\n")
	// 字符串遍历方式2： for-rnage
	for index, value := range mystr {
		fmt.Printf("index=%d, value=%c\n", index, value)
	}
}

func example() {
	for i := 1; i <= 100; i++ {
		j := 6
		if i%j == 0 {
			fmt.Println(i)
		}
	}
}
