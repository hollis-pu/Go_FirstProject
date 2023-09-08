package main

import (
	"fmt"
	"time"
)

/*
*
  - Description:
    协程的简单示例
    请编写一个程序，完成如下功能:
    1)在主线程(可以理解成进程)中，开启一个 goroutine该协程每隔1秒输出"goroutine:hello,world"
    2)在主线程中也每隔一秒输出"main:hello,golang”输出 10 次后，退出程序
    3)要求主线程和 goroutine 同时执行
  - @Author Hollis
  - @Create 2023/9/8 14:27
*/
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("goroutine:hello,golang")
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	start := time.Now().UnixMilli()
	go test() // 开启了一个协程

	for i := 0; i < 10; i++ {
		fmt.Println("main:hello,golang")
		time.Sleep(time.Millisecond * 500)
	}
	end := time.Now().UnixMilli()
	fmt.Println("程序耗时(ms)：", end-start) // 未加协程前：10101ms  加上协程后：5050ms
}
