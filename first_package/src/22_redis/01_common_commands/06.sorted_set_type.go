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
    sorted set类型的redis操作
  - @Author Hollis
  - @Create 2023/9/12 16:19
*/
func main() {
	// 连接到redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	// 1.ZAdd：将所有指定成员添加到键为key有序集合（sortedset）里面。
	// 添加时可以指定多个分数/成员（score/member）对。
	// 如果指定添加的成员已经是有序集合里面的成员，则会更新改成员的分数（score）并更新到正确的排序位置。
	// 如果key不存在，将会创建一个新的有序集合（sortedset）并将分数/成员（score/member）对添加到有序集合，就像原来存在一个空的有序集合一样。
	// 返回添加到有序集合的成员数量。
	result, _ := rdb.ZAdd(ctx, "scores", redis.Z{Score: 80, Member: "Tom"}, redis.Z{Score: 92, Member: "Jerry"}).Result()
	fmt.Println("result:", result) // 2

	// 2.ZCard：返回key的有序集元素个数。
	sortedSetCount, _ := rdb.ZCard(ctx, "scores").Result()
	fmt.Println("sortedSetCount:", sortedSetCount) // 2

	// 3.ZCount：返回有序集key中，score值在min和max之间(默认包括score值等于min或max)的成员。
	sortedSetSelectCount, _ := rdb.ZCount(ctx, "scores", strconv.Itoa(80), strconv.Itoa(90)).Result()
	fmt.Println("sortedSetSelectCount:", sortedSetSelectCount) // 1

	// 4.ZRange：返回存储在有序集合key中的指定范围的元素。返回的元素可以认为是按得分从最低到最高排列。
	// 如果你需要成员按 score 值递减(从大到小)来排列，请使用 ZRevRange 命令。
	sortedSetRange, _ := rdb.ZRange(ctx, "scores", 0, -1).Result()
	fmt.Println("sortedSetRange:", sortedSetRange) // [Tom Jerry]

	// 5.ZRangeByScore：返回key的有序集合中的分数在min和max之间的所有元素（包括分数等于max或者min的元素）。元素被认为是从低分到高分排序的。
	sortedSetRange, _ = rdb.ZRangeByScore(ctx, "scores", &redis.ZRangeBy{
		Min:    strconv.Itoa(80), // 查询值介于[80,90]之间的成员
		Max:    strconv.Itoa(100),
		Offset: 0,
		Count:  0,
	}).Result()
	fmt.Println("sortedSetRange:", sortedSetRange) // [Tom Jerry]

	// 6.ZScore：返回有序集key中，成员member的score值。
	tomScore, _ := rdb.ZScore(ctx, "scores", "Tom").Result()
	fmt.Println("tomScore:", tomScore) // 80

	// 7.ZRank：返回有序集key中成员member的排名。其中有序集成员按score值递增(从小到大)顺序排列。排名以0为底。
	// 从大到小的顺序请用 ZRevRank。
	rank, _ := rdb.ZRank(ctx, "scores", "Jerry").Result()
	fmt.Println("rank:", rank) // 1  排名从0开始（升序），Jerry排名为1

	// 8.ZRem：移除有序集key中的一个或多个成员，不存在的成员将被忽略。
	// ZRemRangeByScore：移除有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。
	// ZRemRangeByRank：移除有序集合 key 中，指定排名(rank)区间内的所有成员。其中区间分别以下标参数 start 和 stop 指出，包含 start 和 stop 在内。
	rdb.ZAdd(ctx, "scores", redis.Z{Score: 68, Member: "Jane"}) // 添加value="Jane"的成员
	rdb.ZRem(ctx, "scores", "Tom")                              // 移除value="Tom"的成员

	sortedSetRange, _ = rdb.ZRange(ctx, "scores", 0, -1).Result()
	fmt.Println("sortedSetRange:", sortedSetRange) // [Jane Jerry]

	rdb.Expire(ctx, "scores", time.Second)
}
