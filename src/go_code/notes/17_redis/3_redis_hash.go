package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "10.250.0.10:6379")
	if err != nil {
		fmt.Println("redis.Dial Error:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("HMSet", "user01", "name", "hybfkuf", "age", "200")
	if err != nil {
		fmt.Println("conn.Do HMSet Error:", err)
		return
	}

	r, err := redis.Strings(conn.Do("HMGet", "user01", "name", "age")) // 此处为 redis.Strings 不是 redis.String
	if err != nil {
		fmt.Println("conn.Do HMGet Error:", err)
		return
	}
	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}
