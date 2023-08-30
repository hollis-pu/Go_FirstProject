package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
*
  - Description:
    生成随机数
    break的使用
    随机生成 1-100 的一个数，直到生成了 99 这个数，统计一共用了多少次
  - @Author Hollis
  - @Create 2023/8/24 19:55
*/
func main() {
	count := 0
	rand.Seed(time.Now().Unix())
	for {
		count++
		data := rand.Intn(100) + 1
		if data == 99 {
			break // 当data等于99时，退出循环
		}
	}
	fmt.Println("Times:", count)
}
