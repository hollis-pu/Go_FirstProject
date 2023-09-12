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
    string类型的redis操作
  - @Author Hollis
  - @Create 2023/9/12 12:46
*/
func main() {
	//连接到redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	// 1.Set：设置key-value
	rdb.Set(ctx, "name", "Tom", 0) // 设置key-value

	// 2.Get：获取key的value
	val, _ := rdb.Get(ctx, "name").Result() // 通过key获取value
	fmt.Println("val:", val)                // val: Tom

	// 3.GetSet：更新一个key的value，并返回这个key的旧值
	oldVal, _ := rdb.GetSet(ctx, "name", "Jerry").Result()
	fmt.Println("oldVal:", oldVal) // oldVal: Tom。但name的value已经被更新为Jerry

	// 4.SetNX：如果key不存在，则设置这个key的value，否则不做任何操作
	rdb.SetNX(ctx, "name", "Jenny", 0) //  由于key=name已经存在，则不进行任何操作
	rdb.SetNX(ctx, "age", 34, 0)       // key=age原来不存在，则创建并设置age=34

	// 5.MSet：批量设置key的值
	rdb.MSet(ctx, "key1", "value1", "key2", "value2", "key3", "value3")

	// 6.MGet：批量获取key的值
	getMultipleValues := rdb.MGet(ctx, "key1", "key2", "key3")
	fmt.Println("getMultipleValues:", getMultipleValues) // [value1 value2 value3]，返回的是多个value组成的切片。

	// 7. Incr/IncrBy/IncrByFloat：针对一个key的数值进行递增操作
	rdb.Set(ctx, "count", 0, 0) // 设置一个count=1的键值对
	// Incr 函数每次加1
	val1, _ := rdb.Incr(ctx, "count").Result() // 将键为count的value值增加1
	fmt.Println("val1:", val1)                 // 1
	// IncrBy 函数，可以指定每次递增多少
	val1, _ = rdb.IncrBy(ctx, "count", 3).Result() // 将键为count的value值增加3
	fmt.Println("val1:", val1)                     // 4
	// IncrByFloat 函数，可以指定每次递增多少，跟IncrBy的区别是累加的浮点数
	rdb.Set(ctx, "countFloat", 0.0, 0)
	valFloat, _ := rdb.IncrByFloat(ctx, "countFloat", 2.2).Result() // 将键为countFloat的value值增加2.2
	fmt.Println("valFloat1:", valFloat)                             // 2.2

	// 8. Decr/DecrBy：针对一个key的数值进行递减操作
	rdb.Set(ctx, "count", 0, 0) // 设置一个count=1的键值对
	// Incr 函数每次减1
	val1, _ = rdb.Decr(ctx, "count").Result() // 将键为count的value值增减1
	fmt.Println("val1:", val1)                // -1
	// IncrBy 函数，可以指定每次递减多少
	val1, _ = rdb.DecrBy(ctx, "count", 3).Result() // 将键为count的value值增减3
	fmt.Println("val1:", val1)                     // -4
	// 没有DecrByFloat函数

	// 9.Del：删除key操作，支持批量删除
	rdb.Del(ctx, "key1")         // 删除1个key
	rdb.Del(ctx, "key2", "key3") // 批量删除key

	// 10.设置过期时间
	rdb.Expire(ctx, "name", time.Second) // 1秒后，name就过期了
	rdb.Expire(ctx, "count", time.Second)
	rdb.Expire(ctx, "countFloat", time.Second)
}
