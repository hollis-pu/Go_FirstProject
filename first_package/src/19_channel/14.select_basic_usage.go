package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
*
  - Description:
    channel中，select的基本使用
  - @Author Hollis
  - @Create 2023/9/10 13:05
*/

func main() {
	channel1 := make(chan int, 5)
	channel2 := make(chan string, 5)

	go func() {
		for i := 0; i < 5; i++ {
			channel1 <- i
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel2 <- "Hello" + strconv.Itoa(i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	time.Sleep(time.Second * 6) // 防止主线程先执行完，这里sleep几秒
	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	// 问题，在实际开发中，可能我们不好确定什么时候关闭该管道
	// 可以使用select 方式可以解决
label:
	for {
		select {
		case v := <-channel1:
			fmt.Println("Received from channel1:", v) // 注意：这里，如果channel1一直没有关闭，也不会一直阻塞而导致deadlock，会自动到下一个case匹配。
		case v := <-channel2:
			fmt.Println("Received from channel2:", v)
		default: // 前面都不匹配时，执行default语句的内容
			fmt.Println("No channel operation ready")
			break label
		}
	}
}
