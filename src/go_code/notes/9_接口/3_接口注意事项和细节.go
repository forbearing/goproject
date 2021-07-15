package main

import "fmt"

type MyInterface interface {
	Say()
}
type MyInterface2 interface {
	Hello()
}
type Student struct{}

func (stu Student) Say() {
	fmt.Println("Student say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer say()")
}

type Monster struct{}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()")
}
func (m Monster) Say() {
	fmt.Println("Monster Say()")
}

func main() {
	// 1.接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
	var stu Student
	var inter1 MyInterface = stu
	inter1.Say()

	// 5. 只要是自定义数据类型，都可以实现接口，不仅仅是结构体类型。
	var i integer = 10
	var inter2 MyInterface = i
	inter2.Say()

	// 8. 一个接口（比如A接口）可以继承多个别的接口（比如B，C接口），这时如果要实现A接口，也必须将B，C接口的方法也全部实现。
	var monster Monster
	var inter3 MyInterface = monster
	var inter4 MyInterface2 = monster
	inter3.Say()
	inter4.Hello()
}

func notes() {
	/*
		===== 注意事项和细节
		1. 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
		2. 接口中所有的方法都没有方法体，即都没有实现的方法
		3. 在 Golang 中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口
		4. 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例（变量）赋给接口类型
		5. 只要是自定义数据类型，都可以实现接口，不仅仅是结构体类型。
		6. 一个自定义类型可以实现多个接口
		7. Golang 接口中不能有任何变量
		8. 一个接口（比如A接口）可以继承多个别的接口（比如B，C接口），这时如果要实现A接口，
		   也必须将B，C接口的方法也全部实现。
		9. Interface 类型默认是一个指针（引用），如果没有对 interface 初始化就使用，那么会输出 nil
		10. 空接口 interface{} 没人呢任何方法，所以所有类型都实现了空接口
	*/
}
