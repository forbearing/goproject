package main

import (
	"github.com/gomodule/redigo/redis"
)

func main() {
	var pool *redis.Pool
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲连接数
		MaxActive:   0,   // 表示数据库的最大连接数，0表示没有限制
		IdleTimeout: 100, // 连接最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化连接的代码，要连接哪个 ip 的 redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	c := pool.Get() // 从连接池汇总取出一个连接
	pool.Close()    // 关闭连接，一旦关闭连接池，就不能从连接池中取出连接
}

/*
Redis 连接池
通过 Golang 对 Redis 操作，还可以通过 Redis 连接池对 Redis 进行操作，流程如下：
1. 事先初始化一定数量的连接，放入到连接池
2. 当 Go 需要操作 Redis 时，直接从 Redis 连接池取出连接即可
3. 这样可以节省临时获取 Redis 连接的时间，从而提高效率

	var pool *redis.Pool
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	c := pool.Get()

*/
