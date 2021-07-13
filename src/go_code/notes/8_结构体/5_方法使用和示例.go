package main

import "fmt"

func main() {
	struct_method()
	example1()
	example2()
	example3()
}

func notes() {
	/*
		===== 方法 =====
		- 在某些情况下，我们需要声明（定义）方法，比如 Person 结构体：除了一些字段外（年龄，姓名）
		  Person 结构体还需要一些行为，比如：说话，学习等。这时就需要方法才能完成
		- Golang 中的方法是作用在指定的数据类型上的（和指定的数据类型绑定），因此自定义类型，都可以有方法，而不仅仅是 struct

		===== 方法的声明和调用 =====
		type A struct {
			Num int
		}
		func(a A) test() {
			fmt.Println(a.Num)
		}
		- func(a A)test() 表示结构体有一个方法，方法名为 test
		- (a A) 体现 test 方法是和 A 类型绑定的
	*/
}

type Person struct {
	Name string
}

// 给 Person 类型绑定一方法
func (p Person) printName() {
	fmt.Println("调用方法:", p.Name)
}
func (p Person) speak() {
	fmt.Printf("%v 是好人\n", p.Name)
}
func (p Person) calc() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func (p Person) calc2(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func (p Person) getSum(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}

func struct_method() {
	// printName 方法和 Person 类型绑定
	// printName 方法只能通过 Person 类型的变量来调用，而不能直接调用，也不能使用其它类型的变量来调用。
	p1 := Person{"hybfkuf"}
	p1.printName() // 调用方法
}
func example1() {
	// 给 Person 结构体添加 speak 方法，输出 xxx 是一个好人
	p2 := Person{"Jonas"}
	p2.speak()
}
func example2() {
	// 给 Person 结构体添加 calc 方法，可以计算 1+...1000 的结果
	p3 := Person{"hello"}
	p3.calc()
}
func example3() {
	// 给 Person 结构体添加 calc2 方法，该方法可以接受一个数n，计算从 1+...+n 的结果
	p4 := Person{"hello"}
	p4.calc2(100)
}
func example4() {
	// 给 Person 结构体添加 getSum 方法，可以计算两个数的和，并返回结果
	p5 := Person{"hello"}
	p5.getSum(7, 8)
}
