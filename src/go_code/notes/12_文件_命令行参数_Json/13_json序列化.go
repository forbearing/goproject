package main

import (
	"encoding/json"
	"fmt"
)

// 演示将结构体、map、切片进行序列化
// 定义一个结构体
type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"monster_age"`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2001-1-1",
		Sal:      80000.0,
		Skill:    "吃饭",
	}
	// 将 monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的解雇
	fmt.Printf("struct 序列化后:\n%v\n", string(data))
}

// 将 map 序列化
func testMap() {
	// 定义一个 map
	var a map[string]interface{}
	// 使用 map 之前需要 make
	a = make(map[string]interface{})
	a["name"] = "hybfkuf"
	a["age"] = 10
	a["address"] = "Shanghai"
	// 将 a 这个 map 进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("map 序列化后:\n%v\n", string(data))
}

// 对切片进行序列化，切片类型是 []map[string]interface{}
func testSlice() {
	var slice []map[string]interface{}

	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "user-1"
	m1["passwd"] = "passwd-1"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "user-2"
	m2["passwd"] = "passwd-2"
	slice = append(slice, m2)

	var m3 map[string]interface{}
	m3 = make(map[string]interface{})
	m3["name"] = "user-3"
	m3["passwd"] = "passwd-4"
	slice = append(slice, m3)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("slice 序列化后:\n%v\n", string(data))
}

// 对基本类型序列化
func testfloat64() {
	num1 := 12.34

	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("float64 序列化后:\n%v\n", string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testfloat64()
}

/*
JSON 用于机器解析和生成，其有效的提升网络传输效率，通常程序在网络传输时会将数据（结构体、map 等）序列化
成 JSON 字符串，到接收方得到 JSON 字符串时，再反序列化成原来的数据类型（结构体，map 等）。这种方式已经
成为各个语言的标准。
*/
