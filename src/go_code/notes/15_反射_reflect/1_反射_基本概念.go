package main

func main() {

}

/*
1.使用反射机制，编写函数的适配器
2.变量、interface{}、reflect.Value 是可以相互转换的
	1.将 interface{} --> reflect.Value		rVal := reflect.ValueOf(b)
	2.将 reflect.Value --> interface{}		iVal := rVal.interface{}
	3.interface{} --> 原来的变量			  使用类型断言


import "reflect"
	reflect 实现了运行时反射，允许程序操作任意类型的对象，典型用法是用静态类型 interface{}
	保存一个值，通过调用 TypeOf 获取其动态类型信息，该函数返回一个 Type 类型值。调用 ValueOf
	函数返回一个 Value 类型值，该值代表运行时的数据。Zero 接收一个 Type 类型参数并返回一个代表
	该类型零值的 Value 类型值。

反射基本介绍：
	1.反射可以在运行时动态获取变量的各种信息，比如变量的类型(type)、类别(kind)。
	2.如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
	3.通过反射，可以修改变量的值，可以调用关联的方法
	4.使用反射，需要 import ("reflect")
反射重要的函数和概念
	1.reflect.TypeOf(变量名), 获取变量的类型，返回 reflect.Type 类型
	2.reflect.ValueOf(变量名), 获取变量的值，返回 reflect.Value 类型
	  reflect.Value 是一个结构体类型。通过 reflect.Value 可以获取该变量很多信息。
反射注意事项和细节说明
	1.reflect.Value.Kind，获取变量的类别，返回的是一个常量
	2.Type 是类型，kind 的是类别，Type 和 kind 可能相同，也可能是不同的
		var num int = 10, num 的 Type 是 int，Kind 也是 int
		var stu Student, stu 的 Type 是 包名.Student, Kind 是 struct
	3.通过反射可以让变量在 interface{} 和 reflect.Value 之间相互转换
	4.通过反射的方式来获取变量的值(并返回对应的类型)，要求数据类型匹配，比如 x 是 int，
	  那么就应该使用 reflect.Value(x).Int()，而不能使用其它的。否则就报 panic。
	5.通过反射来修改变量，注意当使用 SetXxx 方法来设置需要通过对应的指针类型来完成，这样才能
	  改变传入的变量的值，同时需要使用到 reflect.Value.Elem() 方法
*/
