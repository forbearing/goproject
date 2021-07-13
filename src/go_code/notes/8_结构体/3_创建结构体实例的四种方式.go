package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 方式一：
	var p1 Person
	fmt.Println(p1)

	// 方式二：
	p2 := Person{}
	p2.Name = "方式二-1"
	fmt.Println(p2)
	p3 := Person{"方式二-2", 18}
	fmt.Println(p3)

	// 方式三：
	p4 := new(Person) // var p4 *Person = new(Person)
	p4.Name = "方式三"
	fmt.Println(*p4)

	// 方式四：
	p5 := &Person{} //var p5 *Person = &Person{}
	// (*p5).Name <==> p5.Name
	// go 的设计者为了程序使用方便，底层会对 p5.Name = "hybfkuf" 进行处理，会给 p5 加上取值运算 (*p5).Name
	(*p5).Name = "方式四-1"
	p5.Name = "方式四-1"
	fmt.Println(*p5)
	p6 := &Person{"方式四-2", 22} // var p6 *Person = &Person{"Jonas", 22}
	fmt.Println(*p6)
}

func notes() {
	/*
		===== 创建结构体变量和访问结构体字段 =====
		方式一：直接声明
			var person Person
		方式二：{}
			var person Person = Person{}
		方式三：&
			var person *Person = new(Person)
		方式四：{}
			var person *Person = &Person{}

		- 第三种方式和第四种方式返回的是结构体指针
		- 结构体指针访问字段的标准方式应该是 (*结构体指针).字段名，比如 (*person).Name = "Tom"
		  但 Go 做了一个简化，也支持 "结构体指针.字段名"，比如 person.Name = "Tom"，更加符合程序员
		  的习惯。Go 编译器对 perosn.Name 做了转化 (*person).Name

	*/
}
