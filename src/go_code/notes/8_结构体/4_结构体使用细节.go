package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	struc_details()
}

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

func notes() {
	/*
		==== 结构体的注意事项和使用细节 =====
		1. 结构体的所有字段在内存中是连续的
		2. 结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段（名字、个数和类型）
		3. 结构体进行 type 重新定义，相当于取别名，Golang 认为是新的数据类型，但是相互间可以强转
		4. struct 的每个字段上，可以写一个 tag，该 tag 可以通过反映射机制获取，常见的使用场景就是序列化和反序列化
	*/
}

type A struct {
	Num int
}
type B struct {
	Num int
}

type Student struct {
	Name string
	Age  int
}
type Stu Student

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func struc_details() {
	// 1. 结构体的所有字段在内存中是连续的
	// 		r1 有4个 int，在内存中是连续分布的
	//		r2 有两个 *Point 类型，这两个 *Point 类型的本身地址也是连续的，但是他们指向的地址不一定是连续的
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Println(r1.leftUp.x, r1.leftUp.y, r1.rightDown.x, r1.rightDown.y)
	fmt.Println(&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	// 2. 结构体转换
	var a A
	var b B
	b.Num = 10
	a = A(b) // 可以转换，但要求：结构体的字段要完全一样（包括名字、个数、类型）
	fmt.Println(a, b)

	// 3. 结构体重新定义，转换
	var stu1 Student
	var stu2 Stu
	stu2.Name = "hybfkuf"
	stu1 = Student(stu2)
	fmt.Println(stu1, stu2)

	// 4. tag
	//		将 monster 变量序列化为 json 格式字符串
	//		json.Marshal 函数中使用反射
	monster := Monster{"牛魔王", 500, "芭蕉扇"}
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误", err)
	}
	fmt.Println(string(jsonStr))
}
