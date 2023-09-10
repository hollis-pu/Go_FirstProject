package main

import "fmt"

/*
*
  - Description:
    管道容量的使用说明
  - @Author Hollis
  - @Create 2023/9/10 14:12
*/
func main() {
	// 1.无缓冲管道
	// - 当一个goroutine尝试向无缓冲通道发送数据时，它会等待另一个goroutine准备好从通道接收数据，只有当接收操作准备好时，发送操作才会成功。
	ch1 := make(chan int) // 创建一个无缓冲的管道

	fmt.Println("capacity of ch1:", cap(ch1)) // 0
	fmt.Println("len of ch1:", len(ch1))      // 0

	//ch1 <- 1 // 直接向ch1中写入数据会报错：fatal error: all goroutines are asleep - deadlock!

	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()
	num1 := <-ch1
	num2 := <-ch1

	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)

	// 2.有缓冲管道
	// - 当一个goroutine尝试向有缓冲通道发送数据时，如果通道的缓冲区未满，发送操作会成功，否则它会等待直到有足够的空间。
	// - 当一个goroutine尝试从有缓冲通道接收数据时，如果通道的缓冲区不为空，接收操作会成功，否则它会等待直到有数据可用。
	// - 有缓冲通道用于实现异步通信，可以在一定程度上平衡发送和接收的速度，以提高程序的性能。
	ch2 := make(chan int, 3)
	go func() {
		ch2 <- 1
		ch2 <- 2
		ch2 <- 3
	}()

	num1 = <-ch2
	num2 = <-ch2

	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)
}
