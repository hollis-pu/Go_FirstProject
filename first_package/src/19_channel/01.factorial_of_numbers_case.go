package main

import (
	"fmt"
	"time"
)

/*
*
  - Description:
    channel的简单使用。
    入门案例：
    需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。最后显示出来。要求使用goroutine完成
    分析思路：
    1) 使用goroutine 来完成，效率高，但是会出现并发/并行安全问题；
    2) 这里就提出了不同goroutine如何通信的问题。
  - @Author Hollis
  - @Create 2023/9/8 16:30
*/
var (
	myMap01 = make(map[int]int, 10)
)

func main() {
	for i := 1; i <= 200; i++ {
		go test01(i)
	}

	time.Sleep(time.Second)

	for i, v := range myMap01 {
		fmt.Printf("map01[%d]=%d\n", i, v)
	}
}
func test01(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 存在资源竞争的问题！
	myMap01[n] = res // 多个协程同时（并发）向同一个map中写入数据，报错：fatal error: concurrent map writes
}
