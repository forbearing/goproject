package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"monster_age"`
	Birthday string
	Sal      float64
	Skill    string
}

// 将 json 字符串反序列化成 struct
func unmarshalStruct() {
	// 说明：str 在项目开发中，是通过网络传输获取到的，或者是读取文件获取到的。
	str := "{\"monster_name\":\"牛魔王\",\"monster_age\":500,\"Birthday\":\"2001-1-1\",\"Sal\":80000,\"Skill\":\"吃饭\"}"
	// 定义一个 monster 实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal error = %v\n", err)
	}
	fmt.Printf("反序列化 struct: \n%v\n", monster)
}

// 将 json 字符串反序列化成 map
func unmarshalMap() {
	str := "{\"address\":\"Shanghai\",\"age\":10,\"name\":\"hybfkuf\"}"
	// 定义一个 map
	// 反序列化 map，不需要make，因为 make 操作封装到 Unmarshal 函数
	var a map[string]interface{}
	//反序列化
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal error = %v\n", err)
	}
	fmt.Printf("反序列化 map: \n%v\n", a)
}

func unmarshalSlice() {
	str := "[{\"name\":\"user-1\",\"passwd\":\"passwd-1\"}," +
		"{\"name\":\"user-2\",\"passwd\":\"passwd-2\"}," +
		"{\"name\":\"user-3\",\"passwd\":\"passwd-4\"}]"
	// 定义一个 slice
	var slice []map[string]interface{}
	// 反序列化，不需要 make，因为 make操作被封装到 unmarshal 函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal error = %v\n", err)
	}
	fmt.Printf("反序列化 slice: \n%v\n", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}

// json 反序列化：将 json 字符串反序列化成对应的数据类型（比如结构体
/*
	1. 在反序列化一个 Json 字符串，要确保反序列化后的数据类型和原来序列化前的数据类型一致。
	2. 如果 json 字符串是通过程序获取的，则不需要再对 " 转义处理
*/
