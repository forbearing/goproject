package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64        // 数组
	ptr    *int              // 指针
	slice  []int             // 切片
	map1   map[string]string // map
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	// 结构体字段如果是slice、map 需要先 make
	var p1 Person
	fmt.Println(p1)
	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	p1.slice = make([]int, 3)
	p1.slice[0] = 99
	fmt.Println(p1)

	p1.map1 = make(map[string]string)
	p1.map1["name"] = "hybfkuf"
	p1.map1["age"] = "18"
	fmt.Println(p1)

	// 不同结构体变量的字段是独立，互不影响，一个结构体字段的更改不影响其它的字段，
	// 结构体是值类型
	fmt.Println()
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500
	monster2 := monster1
	fmt.Println("monster1 =", monster1)
	fmt.Println("monster2 =", monster2)
	monster2.Name = "孙悟空"
	fmt.Println("monster1 =", monster1)
	fmt.Println("monster2 =", monster2)
}

func notes() {
	/*
		===== 如何声明结构体 =====
		type 结构体名称 struct {
			field1 type
			field2 type
		}
		===== 基本介绍：=====
		1. 从概念或叫法上：结构体字段 = 属性 = field
		2. 字段是结构体的一个组成部分，一般是基本数据类型、数组、也可以是引用类型。

		===== 字段和属性注意事项和细节： =====
		1. 字段声明语法如同变量，示例：字段名 字段类型
		2. 字段的类型可以为：基本类型、数组或引用类型
		3. 在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值（默认值）
			布尔默认 false，整型是0，字符串是""
			数组类型的默认值和它的元素类型相关，score[3]int 为 [0,0,0]
			指针、slice、map 的零值为 nil，即还没有分配空间
		4. 不同结构体变量的字段是独立的，互不影响，一个结构体变量字段的更改，不影响另外一个。
		5. 结构体是值类型
	*/
}
