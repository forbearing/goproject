package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// 定义一个全局的 pool
var pool *redis.Pool

// 当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "10.250.0.10:6379")
		},
	}
}

func main() {
	// 先从 pool 取出一个连接
	// 如果要从 pool 取出连接，一定要保证连接池是没有关闭的
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", "name", "hybfkuf~~~")
	if err != nil {
		fmt.Println("conn.Do set error:", err)
		return
	}
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("conn.Do get error:", err)
		return
	}
	fmt.Println("r=", r)
	time.Sleep(time.Second * 10000)
}
