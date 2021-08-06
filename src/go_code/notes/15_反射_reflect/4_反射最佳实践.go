package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"monster_name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"monster_score"`
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("----- start -----")
	fmt.Println(s)
	fmt.Println("----- end -----")
}
func (s Monster) GetSum(n1 int, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func testStruct(a interface{}) {
	rTyp := reflect.TypeOf(a)  // get reflect.Type
	rVal := reflect.ValueOf(a) // get reflect.Value
	kd := rVal.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	// 获取到该结构体有几个字段
	num := rVal.NumField()
	fmt.Printf("struct have %d fields\n", num)
	// 遍历结构体变量所有的字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: Value=%v\n", i, rVal.Field(i))
		// 获取 struct 标签，注意需要通过 reflect.Type 来获取 tag 标签的值
		tagVal := rTyp.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: Tag=%v\n", i, tagVal)
		}
	}

	// 获取到该结构体有多少个方法
	numOfMethod := rVal.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	// 调用结构体第二个方法。第二个方法无参数，（方法的排序默认是按照函数名的排序）
	rVal.Method(1).Call(nil)
	// 调用接哦固体一个方法。第一个方法有参数
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rVal.Method(0).Call(params) // 传入的是 []reflect.Value
	fmt.Println("res=", res[0].Int())  // 返回结果也是 []reflect.Value

}

func main() {
	// 使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值。
	var a Monster = Monster{
		Name:  "hybfkuf",
		Age:   20,
		Score: 99.9,
	}
	testStruct(a)

}

/*
反射最佳实践
	1. 使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值。
	2. 使用到的方法
		rVal.NumField()
		rVal.Field()
		rTyp.Field(i).Tag.Get("json")
		rVal.NumMethod()
		rVal.Method()
		rVal.Method(0).Call(params)
		rVal.
*/
