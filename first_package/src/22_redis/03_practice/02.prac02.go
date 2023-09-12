package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

/*
*
* Description:
要求
1.student信息[name, age, gender]
2.通过终端输入三个student的信息，使用golang操作redis,存放到redis中[比如使用hash数据类型]
3.编程，遍历出所有的student信息，并显示在终端
4.提示，保存student可以使用 hash数据类型， 遍历时先取出所有的keys
* @Author Hollis
* @Create 2023/9/12 20:34
*/
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	rdb.HMSet(ctx, "student", "name", "tom", "age", 28, "gender", "male")

	stuMap, _ := rdb.HGetAll(ctx, "student").Result()
	for k, v := range stuMap {
		fmt.Printf("k:%v v:%v\n", k, v)
	}
	rdb.Expire(ctx, "student", time.Second)
}
