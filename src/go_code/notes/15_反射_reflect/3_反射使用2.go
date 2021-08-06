package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改 num int 的值，修改 Student 的值

func reflect01(b interface{}) {
	// 1. 获取 reflect.Value
	rVal := reflect.ValueOf(b)
	// 传入的是一个指针，此时 rVal 也是指针类型
	fmt.Printf("rVal kind=%v", rVal.Kind())

	rVal.Elem().SetInt(20)
}

func reflect02(a interface{}) {
	rTyp := reflect.TypeOf(a)
	rVal := reflect.ValueOf(a)

	fmt.Println(rTyp.Kind(), rVal.Kind())

	f2 := rVal.Interface()
	f3, ok := f2.(float64)
	if ok {
		fmt.Println(f3)
	}

	rVal.Elem().SetFloat(2.2)
}

func main() {
	num := 10
	// 不能使用值拷贝，需要引用拷贝
	// reflect01(num)
	fmt.Println(num)
	reflect01(&num)
	fmt.Println(num)
	fmt.Println()

	f1 := 1.1
	fmt.Println(f1)
	reflect02(&f1)
	fmt.Println(f1)

}
