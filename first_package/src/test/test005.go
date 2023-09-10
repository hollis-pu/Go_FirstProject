package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023/9/10 13:34
 */
func main() {
	ch1 := make(chan int, 1)

	fmt.Println("capacity of ch1:", cap(ch1)) // 0
	fmt.Println("len of ch1:", len(ch1))      // 0

	//ch1 <- 1 // 直接向ch1中写入数据会报错：fatal error: all goroutines are asleep - deadlock!

	go func() { // 在协程中使用ch1时，可以认为其容量是无限大的。
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()
	num1 := <-ch1
	num2 := <-ch1
	fmt.Println("capacity of ch1:", cap(ch1)) // 其实际容量还是0，但可以存入数据。
	fmt.Println("len of ch1:", len(ch1))      // 0
	fmt.Println("num:", num1)                 // 1
	fmt.Println("num:", num2)                 // 2

}
