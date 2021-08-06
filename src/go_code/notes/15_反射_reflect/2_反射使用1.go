package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	// 通过反射获取传入的变量的 type, kind, value

	// 1.先获取 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Printf("rType=%v, rValue type %T\n", rTyp, rTyp)

	// 2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rValue=%v, rValue type %T\n", rVal, rVal)
	n2 := 2 + rVal.Int()
	fmt.Println(n2)

	// 3.将 rVal 转成 interface{}
	iV := rVal.Interface()

	// 4.将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2=", num2)
}

func reflectTest02(b interface{}) {
	// 1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rValue=%v, rValue type %T\n", rVal, rVal)
	// (1) 获取到 kind
	kind1 := rTyp.Kind()
	kind2 := rVal.Kind()
	fmt.Printf("kind=%v kind=%v\n", kind1, kind2)

	// 3. 将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV=%v, iV type %T\n", iV, iV)

	// 4. 将 interface{} 通过断言转成需要的类型
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {
	// 演示(基本数据类型、interface{}、reflect.Value) 进行反射的基本操作

	// 1.定义一个 int
	// num := 10
	// reflectTest01(num)

	// 2.定义一个 Student 结构体
	stu := Student{Name: "hybfkuf", Age: 20}
	reflectTest02(stu)

}
