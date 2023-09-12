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
    hash类型的redis操作
  - @Author Hollis
  - @Create 2023/9/12 13:45
*/
func main() {
	//连接到redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	// 1.HSet：设置hash字段值
	rdb.HSet(ctx, "user_1", "username", "tom") // 分别设置key，field和value

	// 2.HGet：获取hash字段值
	hashValue, _ := rdb.HGet(ctx, "user_1", "username").Result() //  通过key和field，获取value
	fmt.Println("hashValue:", hashValue)                         // hashValue: tom

	// 3.HSetNX：如果field字段不存在，则设置hash字段值。否则不进行任何操作
	rdb.HSetNX(ctx, "user_1", "username", "jerry") // key=user_1,field=username字段已经存在，所以不进行任何操作
	rdb.HSetNX(ctx, "user_2", "username", "jerry") // key=user_2,field=username字段不存在，对字段值进行设置。

	// 4.HMSet：设置hash一个key的多个field和value
	rdb.HMSet(ctx, "person", "name", "tom", "age", 26, "job", "doctor")

	// 5.HMGet：获取hash一个key的多个field和value
	getMultipleFiled := rdb.HMGet(ctx, "person", "name", "age", "job")
	fmt.Println("getMultipleFiled:", getMultipleFiled) // hmget person name age job: [tom 26 doctor]

	// 6.HGetAll：获取hash一个key的所有field和value
	getAllFiled, _ := rdb.HGetAll(ctx, "person").Result() // 这里的返回值getAllFiled是一个map类型
	fmt.Println("getAllFiled:", getAllFiled)              // map[age:26 job:doctor name:tom]

	// 7.HIncrBy/HIncrByFloat：根据key和field字段，对value的值进行累加
	rdb.HSet(ctx, "book", "count", 0)
	count, _ := rdb.HIncrBy(ctx, "book", "count", 1).Result()
	fmt.Println("count:", count) // 1

	// 8.HKeys：根据key，返回所有的field
	allFields, _ := rdb.HKeys(ctx, "person").Result() // 这里返回的是一个切片
	fmt.Println("allFields:", allFields)              // [name age job]

	// 9.HLen：根据key，查询hash的field数量
	fieldLen, _ := rdb.HLen(ctx, "person").Result()
	fmt.Println("fieldLen:", fieldLen) // 3

	// 10.HDel：根据key和field，删除hash字段，支持批量删除hash字段（多个field）
	rdb.HDel(ctx, "person", "age", "job") // key=person的hash表中，就只剩下了name字段

	// 11.HExists：检测hash字段名是否存在
	isAgeExists, _ := rdb.HExists(ctx, "person", "age").Result()
	isNameExists, _ := rdb.HExists(ctx, "person", "name").Result()
	fmt.Println("isAgeExists:", isAgeExists)   // false，age字段已经被删除了
	fmt.Println("isNameExists:", isNameExists) // true，name字段还保留着的

	rdb.Expire(ctx, "person", time.Second)
	rdb.Expire(ctx, "user_1", time.Second)
	rdb.Expire(ctx, "user_2", time.Second)
	rdb.Expire(ctx, "book", time.Second)
}
