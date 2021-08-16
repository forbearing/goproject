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

	_, err = conn.Do("HSet", "user01", "name", "hybfkuf")
	if err != nil {
		fmt.Println("conn.Do HSet Error:", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", "200")
	if err != nil {
		fmt.Println("conn.Do HSet Error:", err)
		return
	}

	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("conn.Do HGet Error:", err)
		return
	}
	fmt.Println("name", r1)

	r2, err := redis.String(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("conn.Do HGet Error:", err)
		return
	}
	fmt.Println("age:", r2)

}
