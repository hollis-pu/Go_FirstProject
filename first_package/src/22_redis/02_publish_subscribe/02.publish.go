package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

/*
*
  - Description:
    发布channel
  - @Author Hollis
  - @Create 2023/9/12 19:17
*/
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	// 发布channel
	rdb.Publish(ctx, "channel1", "hello world")

	// 查询订阅者信息
	chs, _ := rdb.PubSubNumSub(ctx, "channel1").Result()
	for ch, count := range chs { // ch：channel名称，count：订阅者数量
		fmt.Printf("%s subscribers count:%v\n", ch, count)
	}
}
