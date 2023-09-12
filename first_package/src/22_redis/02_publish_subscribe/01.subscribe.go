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
    订阅channel
  - @Author Hollis
  - @Create 2023/9/12 19:12
*/
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	sub := rdb.Subscribe(ctx, "channel1") // 订阅channel1
	// PSubscribe 用法跟Subscribe一样，区别是PSubscribe订阅通道(channel)支持模式匹配。如：rdb.Subscribe(ctx, "channel_*")，*代表任意字符。
	// sub.Channel() 返回go channel，可以循环读取redis服务器发过来的消息
	//for msg := range sub.Channel() {
	//	// 打印收到的消息
	//	fmt.Printf("Channel:%v msg:%v\n", msg.Channel, msg.Payload)
	//}
	for {
		message, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %v\tmsg:%v\n", time.Now().Format("2006-01-02 15:04:05"), message.Channel, message.Payload)
	}

	//sub.Unsubscribe(ctx, "channel1")	//取消订阅
}
