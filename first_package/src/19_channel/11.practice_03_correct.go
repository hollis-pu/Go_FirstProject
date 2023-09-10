package main

import (
	"fmt"
	"time"
)

/*
*
  - Description:
    对10的改进。求出小于数字n的所有素数（使用goroutine）
  - @Author Hollis
  - @Create 2023/9/9 16:14
*/
//

func main() {
	start := time.Now().UnixMilli() // 返回毫秒数
	var (
		n          = 100000
		numChan    = make(chan int, 100) // 存放原始数据的管道
		resultChan = make(chan int, 100) // 存放结果的管道。注意，这个结果管道的容量可以是任意的，而不是必须要比结果集更大，因为管道的内容可以边读边写。
		dataSlice  = make([]int, 0)      // 存放结果的切片
		//writeFinish  = make(chan bool, 1)
		exitChan = make(chan bool, 8) // 判断n个协程是否都已经完成
	)

	go writeDataToChan02(n, numChan)
	for i := 0; i < 8; i++ { // 如果是多核CPU，则会将这些协程分配到不同的逻辑CPU上。不如电脑的核心数8，则开启8个协程可以使CPU的利用率达到最大，再往上增加协程数量，性能提升就不是很明显了。
		go handleData02(n, numChan, resultChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ { // 这里判断n个协程是否都已经完成了
			<-exitChan
		}
		close(resultChan)
	}()

	count := 0
	for result := range resultChan { // 读出管道中的数据，存入到切片中
		dataSlice = append(dataSlice, result)
		count++
	}

	fmt.Println(dataSlice)
	fmt.Println("count:", count)

	end := time.Now().UnixMilli() // 返回毫秒数
	fmt.Println("多协程的方式执行用时(ms)：", end-start)
	// 4个handleData02管道：1021ms
	// 100个handleData02管道：636ms
	// 没有用goroutine和channel的用时：1874ms

}

func writeDataToChan02(n int, numChan chan int) {
	for i := 0; i < n; i++ {
		numChan <- i
	}
	close(numChan)
	//writeFinish <- true  // 这里不用让numChan管道写完，再处理数据，可以一边写入，一边读出处理。所以这里不用加writeFinish控制两个协程的顺序。
}

func handleData02(n int, numChan chan int, resultChan chan int, exitChan chan bool) {
	//<-writeFinish
	for i := 0; i < n; i++ {
		num, ok := <-numChan
		if !ok {
			break
		}
		if isPrime02(num) {
			//dataSlice = append(dataSlice, num)  // 这里存在多个协程向切片中写入数据的资源竞争问题，改成存入到管道中。
			resultChan <- num
		}
	}
	//close(resultChan) // 这里不能直接关闭管道，因为可能其他的协程还为完成工作。这里只需要向exitChan管道写入一个数据即可。
	exitChan <- true // 每完成1个协程，就向exitChan管道写入一个值
}

func isPrime02(num int) bool {
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
