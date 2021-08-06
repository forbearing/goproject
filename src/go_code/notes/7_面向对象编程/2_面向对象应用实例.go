package main

import "fmt"

// 案例一：
// 1、编写一个 Student 的结构体，包含 name、gender、age、id、score 字段，分别为 string、string、int、int、float64
// 2、结构体中声明一个 say 方法，返回 string 类型，方法返回信息中包含所有字段值。
// 3、在 mian 方法中，创建 Student 结构体实例（变量），并访问 say 方法，并将结果打印输出
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("Student 的信息是 name=[%v] gender=[%v] age[%v] id=[%v] score=[%v]",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

// 案例二：
// 1、创建一个 Box 结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高都要从终端输入
// 2、声明一个方法获取立方体的提体积
// 3、创建一个 Box 结构体变量，打印给定尺寸的立方体的体积
type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolume() float64 {
	return box.len * box.width * box.height
}

func main() {
	// 案例一：
	var student = Student{
		name:   "hybfkuf",
		gender: "male",
		age:    18,
		id:     1000,
		score:  89,
	}
	str := student.say()
	fmt.Println(str)

	// 案例二：
	var box Box
	box.len = 1.1
	box.width = 2.2
	box.height = 3.3
	volume := box.getVolume()
	fmt.Println("volume =", volume)

}

/*
步骤
1、声明/定义结构体，确定结构体名
2、编写结构体的字段
3、编写结构体的方法
*/
