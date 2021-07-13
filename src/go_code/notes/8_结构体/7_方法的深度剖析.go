package main

import "fmt"

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	c.radius += 1
	return 3.14 * c.radius * c.radius
}

type integer int

func (i integer) print() {
	fmt.Printf("i = %v\n", i)
}
func (i *integer) change() {
	*i = *i + 1
}

type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main() {
	myCircle := Circle{4.0}
	// sum := (&myCircle).area()
	// 编译器底层做了优化 (&myCircle).area() 等价于 c.area()
	fmt.Printf("before myCircle.radius=%v\n", myCircle.radius)
	sum := myCircle.area()
	fmt.Printf("after myCircle.radius=%v\n", myCircle.radius)
	fmt.Println(sum)

	// Golang 中的方法作用在指定的数据类型上
	var i integer = 10
	i.print()
	i.change()
	fmt.Println(i)

	// 如果实现了 *Student 类型的 String 方法，就会自动调用
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(&stu)
}

func notes() {
	/*
		===== 方法的声明（定义）=====
		func (recevier type) methodName(参数列表) (返回值列表) {
			方法体
			return 返回值
		}
		1. 参数列表：表示方法输入
		2. receiver type 表示这个方法和 type 这个类型锁定，或者说该方法作用于 type 类型
		3. receiver type：type 可以是结构体，也可以是自定义类型
		4. receiver 就是 type 类型的一个变量（实例），比如：Person 结构体的一个变量（实例）
		5. 参数列表：表示方法输入
		6. 返回值列表：表示返回的值，可以是多个
		7. 方法主体：表示为了实现某一功能代码块
		8. return 语句不是必须的

		===== 方法的注意事项和细节讨论 =====
		1. 结构体类型是值类型，在方法调用中，遵守值类型的的传递机制，是值拷贝传递方法
		2. 如程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理
		3. Golang 中的方法作用在指定的数据类型上（即：和指定的类型绑定）因此自定义类型，都有方法。
		   而不仅仅是 struct，比如 int，float32 等都可以有方法
		4. 方法的访问范围控制的规则，和函数一样。方法名首字母小写，只能在本包访问，方法名首字母大写，可以在本包和其它访问
		5. 如果一个变量实现了 String() 方法，那么 fmt.Println 默认会调用这个变量的 String() 进行输出
	*/
}
