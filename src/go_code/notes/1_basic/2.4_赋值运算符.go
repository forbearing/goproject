package main

import "fmt"

func main() {
	my_func()
}

func my_func() {
	/*
		1. 赋值运算符就是将某个运算后的值，赋给指定的变量
			=
			+=
			-=
			*=
			/=
			%=
			<<=		左移后赋值
			>>=		右移后赋值
			&=		按位与后赋值
			^=		按位异或后赋值
			|=		按位或后赋值
		2. 赋值运算符特点
			1. 运算顺序从右往左
			2. 赋值运算符的左边只能是变量，右边可以是变量、表达式、常量值
			3. 比如：a +=3 等价于 a + a + 3
	*/

	// a, b 两个变量的值进行交换
	a := 10
	b := 9
	t := a
	a = b
	b = t
	fmt.Println(a)
	fmt.Println(b)

	i := 10
	i += 7
	fmt.Println((i))
}
