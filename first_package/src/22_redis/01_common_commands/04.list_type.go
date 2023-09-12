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
    list类型的redis操作
  - @Author Hollis
  - @Create 2023/9/12 14:44
*/
func main() {
	// 连接到redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	// 1.LPush：将所有指定的值插入到存于key的列表的头部。如果key不存在，那么在进行push操作前会创建一个空列表。
	rdb.LPush(ctx, "nums", 1).Result()
	// LPush支持一次插入任意个数据
	listLen, _ := rdb.LPush(ctx, "nums", 2, 3, 4, 5, 6).Result() // 返回push操作后的list长度
	fmt.Println("listLen:", listLen)                             // 6

	// 2.LPushX：只有当key已经存在并且存着一个list的时候，在这个key下面的list的头部插入value。与LPush相反，当key不存在的时候不会进行任何操作。
	rdb.LPushX(ctx, "list1", 1) // list1不存在，不会进行任何操作。

	// 3.LPop：移除并且返回key对应的list的第一个元素。
	popElem, _ := rdb.LPop(ctx, "nums").Result()
	fmt.Println("popElem:", popElem) // 6

	// 4.LLen：返回存储在key里的list的长度。如果key不存在，那么就被看作是空list，并且返回长度为0。当存储在key里的值不是一个list的话，会返回error。
	lLen := rdb.LLen(ctx, "nums")
	fmt.Println("lLen:", lLen) // 5

	// 8.LRange：返回存储在key的列表里指定范围内的元素。
	// start和end偏移量都是基于0的下标，即list的第一个元素下标是0（list的表头），第二个元素下标是1，以此类推。
	// 偏移量也可以是负数，表示偏移量是从list尾部开始计数。例如，-1表示列表的最后一个元素，-2是倒数第二个，以此类推。
	listRange, _ := rdb.LRange(ctx, "nums", 0, -1).Result() // 这里返回一个切片
	fmt.Println("listRange:", listRange)                    // [5 4 3 2 1]

	// 9.LRem：从存于key的列表里移除前count次出现的值为value的元素。返回被移除的元素个数。
	//这个count参数通过下面几种方式影响这个操作：
	//	count>0:从头往尾移除值为value的元素。
	//	count<0:从尾往头移除值为value的元素。
	//	count=0:移除所有值为value的元素。
	removeCount, _ := rdb.LRem(ctx, "nums", 3, 1).Result() // 移除前3次出现的value:1
	fmt.Println("removeCount:", removeCount)               // 由于列表中只有1个元素值为1，所以移除次数为1，返回1

	//10.LIndex：返回列表中，下标为 index 的元素（index从0开始）
	index0, _ := rdb.LIndex(ctx, "nums", 0).Result()
	fmt.Println("index0:", index0) // 5

	// 11.LInsert：把 value 插入存于 key 的列表中在基准值 pivot 的前面或后面。返回经过插入操作后的list长度。
	rdb.LInsert(ctx, "nums", "before", 4, "hello") // 在值为4的元素前面插入元素“hello”

	// 遍历list
	listRange, _ = rdb.LRange(ctx, "nums", 0, -1).Result() // 这里返回一个切片
	fmt.Println("listRange:", listRange)                   // [5 hello 4 3 2]

	rdb.Expire(ctx, "nums", time.Second) // 设置nums的过期时间为1秒，防止影响下次的执行结果。
	// 也可以选择将列表清空
	//	listLen, _ = rdb.LLen(ctx, "nums").Result()
	//	for i := 0; i < int(listLen); i++ {
	//		rdb.LPop(ctx, "nums")
	//	}
}
