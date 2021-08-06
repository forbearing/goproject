package main

import "fmt"

type student struct {
	Name  string
	Score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}

func main() {
	var stu = NewStudent("tom", 88.8)
	fmt.Println(stu)

}

/*
Golang 没有构造函数，通常使用工厂模式来解决这个问题
*/
