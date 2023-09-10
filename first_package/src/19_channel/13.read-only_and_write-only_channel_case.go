package main

import (
	"fmt"
	"time"
)

/*
*
  - Description:
    只读和只写管道的案例
  - @Author Hollis
  - @Create 2023/9/10 11:23
*/
func main() {
	var chan01 = make(chan int, 10)
	go writeOnly(chan01)
	go readOnly(chan01)
	time.Sleep(time.Second) // 防止主线程先执行完毕
}

func writeOnly(writeChan chan<- int) { // 在写函数中，将形参的channel设置为只写。这样可以防止误操作，以及底层进行了优化，加快速度。
	for i := 0; i < 100; i++ {
		writeChan <- i
	}
	close(writeChan)
}

func readOnly(readChan <-chan int) { // 在读函数中，将形参的channel设置为只读
	for i := 0; i < 100; i++ {
		num, err := <-readChan
		if !err {
			break
		}
		fmt.Println("num:", num)
	}

}
