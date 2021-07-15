package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// 编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for index, value := range items {
		switch value.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是 %v\n", index, value)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 int/int32/int64 类型，值是 %v\n", index, value)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是 %v\n", index, value)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值是 %v\n", index, value)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是 %v\n", index, value)
		case Student:
			fmt.Printf("第%v个参数是 Student 类型，值是 %v\n", index, value)
		case *Student:
			fmt.Printf("第%v个参数是 *Student 类型，值是 %v\n", index, value)
		default:
			fmt.Printf("第%v个参数是类型不确定，值是 %v\n", index, value)
		}
	}
}

func main() {
	var b1 bool = true
	var i1 int = 1
	var i2 int32 = 2
	var i3 int64 = 3
	var f1 float32 = 1.1
	var f2 float64 = 2.2
	var s1 string = "hello golang"

	stu1 := Student{"hybfkuf", 20}
	stu2 := &Student{"Jonas", 20}

	TypeJudge(b1, i1, i2, i3, f1, f2, s1, stu1, stu2)
}
