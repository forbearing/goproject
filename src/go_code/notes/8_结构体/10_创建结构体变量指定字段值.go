package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (stu Student) Print1() {
	fmt.Printf("Name: %v, Age: %v\n", stu.Name, stu.Age)
}
func (stu *Student) Print2() {
	fmt.Printf("Name: %v, Age: %v\n", stu.Name, stu.Age)
}

func main() {
	var stu1 Student = Student{"Linux-1", 1}
	stu1.Print1()
	stu1.Print2()

	stu2 := Student{"Linux-2", 2}
	stu2.Print1()
	stu2.Print2()

	var stu3 Student = Student{
		Name: "Linux-3",
		Age:  3,
	}
	stu3.Print1()
	stu3.Print2()

	stu4 := Student{
		Name: "Linux-4",
		Age:  4,
	}
	stu4.Print1()
	stu4.Print2()

	var stu5 *Student = &Student{"Linux-5", 5}
	stu5.Print1()
	stu5.Print2()

	stu6 := &Student{
		Name: "Linux-6",
		Age:  6,
	}
	stu6.Print1()
	stu6.Print2()
}

/*
Golang 在创建结构体实例（变量）时，可以直接指定字段的值
方式一：
	var stu Student = Student {"tom", 10}
	var stu := Student{"tom", 10}
	var stu Student = Student {
		Name : "hybfkuf",
		Age : 10
	}
	stu := Student {
		Name : "hybfkuf",
		Age : 10
	}
方式二：
	var stu *Student = &Student{"tom", 10}
	var stu *Student = &Student{
		Name : "hybfkuf",
		Age : 80,
	}
*/
