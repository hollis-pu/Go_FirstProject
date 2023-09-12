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
    set类型的redis操作
  - @Author Hollis
  - @Create 2023/9/12 15:42
*/
func main() {
	// 连接到redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	// 1.SAdd：添加一个或多个指定的member元素到集合的key中，并且自动去重。如果添加的元素已经在集合key中存在则忽略。
	// 如果集合key不存在，则新建集合key，并添加member元素到集合key中。
	// 返回新成功添加到集合里元素的数量。
	sucCount, _ := rdb.SAdd(ctx, "set01", "hello", "world").Result()
	fmt.Println("sucCount:", sucCount) // 2

	sucCount, _ = rdb.SAdd(ctx, "set01", "hello").Result()
	fmt.Println("sucCount:", sucCount) // 再次添加已经存在的元素，添加失败，返回0

	// 2.SCard：返回集合存储的key的基数(集合元素的数量)。
	setCount, _ := rdb.SCard(ctx, "set01").Result()
	fmt.Println("setCount:", setCount) // 2

	// 3.Sismember：返回成员member是否是存储的集合key的成员。
	isMember, _ := rdb.SIsMember(ctx, "set01", "hello").Result()
	fmt.Println("isMember:", isMember) // true

	// 4.Smembers：返回key集合所有的元素。
	setMembers, _ := rdb.SMembers(ctx, "set01").Result() // 这里返回的是一个切片
	fmt.Println("setMembers:", setMembers)               //  [world hello]

	// 5.SRem：在key集合中移除指定的元素。如果指定的元素不是key集合中的元素则忽略。如果key集合不存在则被视为一个空的集合，该命令返回0。
	// 返回从集合中移除元素的个数。
	removeCount, _ := rdb.SRem(ctx, "set01", "hello").Result()
	fmt.Println("removeCount:", removeCount) // 1

	rdb.Expire(ctx, "set01", time.Second)
}
