package main

import "fmt"

/*
*
  - Description:
    遍历channel：使用for+range循环，不能使用fori+len(...)循环
  - @Author Hollis
  - @Create 2023/9/8 21:12
*/
func main() {
	var intChan = make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan <- i * 2 // 放入100个数据到管道中
	}

	// 在遍历时，如果`channel`没有关闭，则会出现`deadlock`的错误，因为`range`循环将会一直等待，从而导致程序陷入死锁状态。
	close(intChan) // 关闭管道后，则可以正常遍历管道
	for v := range intChan {
		fmt.Println("v=", v)
	}
}
