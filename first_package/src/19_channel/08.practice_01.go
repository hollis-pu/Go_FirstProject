package main

/*
*
- Description:
要求:
1.启动一个协程，将1-2000的数放入到一个channel中，比如numChan
2.启动8个协程，从numChan取出数(比如n)，并计算1+...+n的值，并存放到resChan
3.最后8个协程将同完成工作后，再遍历resChan, 显示结果。如 ：res[1]= 1... res[10]= 55...
4.注意:考虑 resChan chan int 是否合适?
- @Author Hollis
- @Create 2023/9/9 09:21
*/

import (
	"fmt"
	"time"
)

var (
	numChan = make(chan int, 2000)
	resChan = make(chan int, 2000)
)

func main() {
	start := time.Now().UnixMilli()

	go putNums()

	for i := 0; i < 8; i++ { // 启8个协程，完成数据计算和写入的工作
		go calcNums()
	}

	i := 1
	for v := range resChan {
		fmt.Printf("v[%d]= %d\n", i, v)
		i++
	}
	end := time.Now().UnixMilli()
	fmt.Println("time(ms):", end-start)

}
func putNums() { // 将数据放入管道
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}
func calcNums() { // 将数据从管道取出并计算，然后存入新的管道
	num := 0
	for num = range numChan {
		sum := 0
		for i := 1; i <= num; i++ {
			sum += i
		}
		resChan <- sum
	}
	if num == 2000 { // 当取到numChan的最后一个元素时，关闭resChan
		close(resChan)
	}
}
