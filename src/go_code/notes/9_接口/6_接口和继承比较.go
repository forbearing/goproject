package main

import "fmt"

// Monkey 结构体
type Monkey struct {
	Name string
}

func (this *Monkey) Climbing() {
	fmt.Println(this.Name, "生来会爬树...")
}

// 声明接口 Bird、Fish
type Bird interface {
	Flying()
}

type Fish interface {
	Swiming()
}

//  LittleMonkey 结构体，让 LittleMonkey 实现 Bird、Fish 接口
type LittleMonkey struct {
	Monkey
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习，会飞翔...")
}
func (this *LittleMonkey) Swiming() {
	fmt.Println(this.Name, "通过学习，会游泳...")
}

func main() {
	monkey := LittleMonkey{Monkey{"悟空"}}
	monkey.Climbing()
	monkey.Flying()
	monkey.Swiming()

}

/*
代码小结：
1. 当 A 结构体继承了 B 结构体，那么 A 结构体就自动继承 B 结构体的字段和方法，并且可以直接使用
2. 当 A 结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可。
	因此接口可以看作是对继承机制的补充。
*/

/*
实现接口 vs 继承
1.接口和继承解决的问题不同
	1.继承的价值主要在于：解决代码的复用性和可维护性
	2.接口的价值主要在于：设计，设计好各种规范（方法），让其他自定义类型去实现这些方法
2.接口比继承更加灵活
	继承是满足 is - a 的关系，而接口只需要满足 like - a 的关系
3.接口在一定程度上实现代码解耦
*/
