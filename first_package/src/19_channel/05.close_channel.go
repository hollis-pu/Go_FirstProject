package main

import "fmt"

/*
*
  - Description:
    关闭管道的基本使用
  - @Author Hollis
  - @Create 2023/9/8 21:05
*/
func main() {
	var intChan = make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) // 关闭管道，关闭后只能读不能写
	//intChan <- 300  // 报错：panic: send on closed channel

	num1 := <-intChan          // 可以从管道中读出数据
	fmt.Println("num1:", num1) // num1: 100
}
