package main

import "fmt"

type FatherInterface interface {
	task01()
}
type MontherInterface interface {
	task02()
}

// 如果需要实现 SonInterface，就需要 FatherInterface 和 MontherInterface 都实现
type SonInterface interface {
	FatherInterface
	MontherInterface
	task03()
}
type Student struct{}

// 空接口, 任何变量都可以赋值给空接口
type T interface{}

func (student Student) task01() {
	fmt.Println("Student doing task01()")
}
func (student Student) task02() {
	fmt.Println("Student doing task02()")
}
func (student Student) task03() {
	fmt.Println("Student doing task03()")
}

func main() {
	// 一个接口（比如A接口）可以继承多个别的接口（比如B，C接口），这时如果要实现A接口，也必须将B，C接口的方法也全部实现。
	var stu1 Student
	var son1 SonInterface = stu1
	son1.task01()
	son1.task02()
	son1.task03()

	var stu2 Student
	var stu3 Student
	// 空接口, 任何变量都可以赋值给空接口
	var t T = stu2
	fmt.Println(t)
	var t2 interface{} = stu3
	fmt.Println(t2)
}

func notes() {

}
