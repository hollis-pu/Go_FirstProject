package main

import "fmt"

/*
*
  - Description:
    求出小于数字n的所有素数（使用goroutine）
  - @Author Hollis
  - @Create 2023/9/9 16:14
*/
// 备注：20230910 这样写出来得到的切片始终为空，但是又找不到问题出在哪里....留到后面再解决吧！
var (
	n            = 100
	numChan001   = make(chan int, n)
	dataSlice    = make([]int, 0)
	writeFinish  = make(chan bool, 1)
	handleFinish = make(chan bool, 1)
)

func main() {
	var (
		n            = 100
		numChan001   = make(chan int, n)
		dataSlice    = make([]int, 0)
		writeFinish  = make(chan bool, 1)
		handleFinish = make(chan bool, 1)
	)

	go writeDataToChan(n, numChan001, writeFinish)
	go handleData(numChan001, dataSlice, handleFinish)
	fmt.Println(dataSlice)
	for {
		_, ok := <-handleFinish
		if !ok {
			break
		}
	}
}

func writeDataToChan(n int, numChan001 chan int, writeFinish chan bool) {
	for i := 0; i < n; i++ {
		numChan001 <- i
	}
	close(numChan001)
	//writeFinish <- true
}

func handleData(numChan001 chan int, dataSlice []int, handleFinish chan bool) {
	//<-writeFinish
	for i := 0; i < n; i++ {
		num, ok := <-numChan001
		if !ok {
			return
		}
		if isPrime(num) {
			dataSlice = append(dataSlice, num)
		}
	}
	handleFinish <- true
	close(handleFinish)
}

func isPrime(num int) bool {
	if num == 2 || num == 3 {
		return true
	}
	if num%2 == 0 || num == 1 {
		return false
	}
	for i := 3; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
