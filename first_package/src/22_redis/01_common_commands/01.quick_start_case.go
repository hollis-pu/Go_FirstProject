package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

/*
*
  - Description:
    redis的快速入门案例
  - @Author Hollis
  - @Create 2023/9/12 9:40
*/
func main() {

	// 连接到redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("rdb:", rdb) // rdb: Redis<localhost:6379 db:0>

	// Set操作
	ctx := context.Background()
	err := rdb.Set(ctx, "name", "tom", 10*time.Second).Err() // 第一个参数：上下文。第二个参数：key。第三个参数：value。第四个参数：过期时间。（0则表示永远不过期）
	if err != nil {
		panic(err)
	}

	// Get操作
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("val:", val) // val: tom

	rdb.Expire(ctx, "name", time.Second)

}
