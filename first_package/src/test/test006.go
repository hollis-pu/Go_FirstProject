package main

import (
	"fmt"
	"time"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/10 13:34
 */
func main() {
	ch1 := make(chan int, 1)

	ch1 <- 1
	ch1 <- 2 // 报错：fatal error: all goroutines are asleep - deadlock!

	//go func() {
	//	ch1 <- 1
	//	ch1 <- 2
	//	ch1 <- 3
	//}()

	time.Sleep(time.Second) // 等待一段时间，确保协程有足够时间触发恐慌

	fmt.Println("Program finished")

}
