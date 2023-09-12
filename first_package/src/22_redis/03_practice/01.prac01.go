package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

/*
*
  - Description:
    Redis的课堂练习
    记录用户浏览商品信息，比如保存商品名。
    编写一个函数，可以取出某个用户最近浏览的10个商品名。
    提示:考虑使用list数据类型
  - @Author Hollis
  - @Create 2023/9/12 20:22
*/
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	for i := 1; i <= 10; i++ {
		rdb.RPush(ctx, "browse_product_records", "product"+strconv.Itoa(i))
	}
	productRecords, _ := rdb.LRange(ctx, "browse_product_records", 0, -1).Result()
	fmt.Println("productRecords:", productRecords)

	rdb.Expire(ctx, "browse_product_records", time.Second*1)
	defer rdb.Close()
}
