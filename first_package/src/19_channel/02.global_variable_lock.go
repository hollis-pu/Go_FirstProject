package main

import (
	"fmt"
	"sync"
	"time"
)

/**
* Description:
	解决01中的资源竞争的问题：
	方式1：全局变量+锁。
	对01中多个协程进行写入操作的map加上锁，每个协程在写的时候lock，写完之后unlock。然后其他协程才能进行写入操作。
* @Author Hollis
* @Create 2023/9/8 17:05
*/

var (
	myMap02 = make(map[int]int, 10)
	lock    sync.Mutex // 全局的互斥锁
)

func main() {
	for i := 1; i <= 200; i++ {
		go test02(i)
	}

	time.Sleep(time.Second * 5) // 这里休眠5秒钟，防止主线程先于协程结束

	lock.Lock() // 读取数据前也要加锁
	for i, v := range myMap02 {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
func test02(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i // 数的阶乘的值会很大，结果会越界，这里将阶乘改为阶加
	}
	// 解决资源竞争的问题
	// 用之前加锁
	lock.Lock()
	myMap02[n] = res
	// 用完之后解锁
	lock.Unlock()
}
