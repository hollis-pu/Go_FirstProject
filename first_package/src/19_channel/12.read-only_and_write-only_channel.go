package main

import "fmt"

/*
*
  - Description:
    只读和只写管道
  - @Author Hollis
  - @Create 2023/9/10 11:12
*/
func main() {
	// 1.默认为可读可写管道
	var chan01 chan int // 可读可写管道
	chan01 = make(chan int, 3)
	fmt.Println(chan01)

	// 2.声明为只写管道
	var chan02 chan<- int // 只写管道
	chan02 = make(chan int, 3)
	chan02 <- 20
	//num := <-chan02 // 报错：invalid operation: cannot receive from send-only channel chan02 (variable of type chan<- int)
	fmt.Println(chan02)

	// 3.声明为只读管道
	var chan03 <-chan int // 只写管道
	chan03 = make(chan int, 3)
	num := <-chan03 // 可读，但由于管道为空，所以报错：fatal error: all goroutines are asleep - deadlock!
	//chan03 <- 20 // 报错：invalid operation: cannot send to receive-only channel chan03 (variable of type <-chan int)
	fmt.Println(chan03)
	fmt.Println(num)
}
