package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// 通过 go 向 redis 写入数据和读取数据

	// 1. 连接到 redis
	conn, err := redis.Dial("tcp", "10.250.0.10:6379")
	if err != nil {
		fmt.Println("redis.Dial Error:", err)
		return
	}
	defer conn.Close()
	// 2. 通过 go 向 redis 写入数据 string [key-val]
	_, err = conn.Do("Set", "from", "来自 golang")
	if err != nil {
		fmt.Println("set.Dial Set error:", err)
		return
	}
	fmt.Println("操纵成功")

	// 3.通过 go 向 redis 读取数据 string [key-val]
	r, err := redis.String(conn.Do("Get", "from")) // r 是 interface{}，需要 redis.String()转换成字符串
	if err != nil {
		fmt.Println("set.Dial Get error", err)
	}
	fmt.Println(r)
}
