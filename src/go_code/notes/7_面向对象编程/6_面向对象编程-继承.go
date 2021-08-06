package main

import "fmt"

// 学生考试系统, student 类
// 将 Pupil 和 Graduate 共有的方法也绑定到 *Student
type Student struct {
	Name  string
	Age   int
	Score float64
}
type Pupil struct {
	Student //嵌入了 Student 匿名结构体
}
type Graduate struct {
	Student // 嵌入了 Student 匿名结构体
}

// 学生类，的方法
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score float64) {
	if score < 0 {
		fmt.Println("分数错误")
		return
	}
	stu.Score = score
}

// 小学生特有方法
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试 ...")
}

// 大学生特有方法
func (g *Graduate) testing() {
	fmt.Println("大学生正在考试 ...")
}

func main() {
	// 当我们对结构体潜入了你名结构体后，使用方法发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "小学生1"
	pupil.Student.Age = 8
	pupil.Student.Score = 0
	pupil.ShowInfo()
	pupil.testing()
	pupil.SetScore(99)
	pupil.ShowInfo()

	fmt.Println()
	graduate := &Graduate{}
	graduate.Student.Name = "大学生1"
	graduate.Student.Age = 21
	graduate.ShowInfo()
	graduate.testing()
	graduate.SetScore(90)
	graduate.ShowInfo()
}

func notes() {
	/*
	   - 继承可以解决代码复用，让我们的编程更加靠近人类思维
	   - 当多个结构体存在相同的属性（字段）和方法时，可以从这个结构体中抽象出结构体，在该结构体中定义这些相同的属性和方法
	   - 其它的结构体不需要重新定义这些属性和方法，只需嵌套一个 Student 匿名结构体即可
	     也就是说，在 Golang 中，如果一个 Student 嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体
	     的字段和方法，从而实现了继承特性。

	   type Goods struct {
	   	Name string
	   	Price int
	   }
	   type Book struct {
	   	Goods		// 这就是嵌套匿名结构体 Goods
	   	Writer string
	   }
	*/

}
