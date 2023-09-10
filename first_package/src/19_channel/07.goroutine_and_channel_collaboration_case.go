package main

import (
	"fmt"
	"time"
)

/*
*
  - Description:
    goroutine和channel协同的案例
    应用实例1：请完成`goroutine`和`channel`协同工作的案例，具体要求：
    1. 开启一个`writeData`协程，向管道`intChan`中写入50个整数。
    2. 开启一个`readData`协程，从管道`intChan`中读取`writeData`写入的数据。
    3. 注意: `writeData`和`readDate`操作的是同一个管道。
    4. 主线程需要等待`writeData`和`readDate`协程都完成工作才能退出。
  - @Author Hollis
  - @Create 2023/9/8 21:48
*/

func main() {
	var (
		intChan  = make(chan int, 50)
		boolChan = make(chan bool, 1)
	)

	go writeData(intChan)
	go readDate(intChan, boolChan)

	for { // 这个for循环的作用是为了防止主线程先于前面的两个协程退出。
		_, ok := <-boolChan
		if !ok {
			break
		}
	}
}

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Printf("写入数据 %d 成功!\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	close(intChan)
}

func readDate(intChan chan int, boolChan chan bool) {
	for i := 0; i < 50; i++ {
		num, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("读取数据 %d 成功！\n", num)
	}
	boolChan <- true
	close(boolChan)
}
