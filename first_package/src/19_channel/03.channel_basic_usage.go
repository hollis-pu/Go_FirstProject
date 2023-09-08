package main

import "fmt"

/*
*
  - Description:
    管道的基本使用
  - @Author Hollis
  - @Create 2023/9/8 20:10
*/
func main() {
	// 1.管道的定义
	var intChan = make(chan int, 3) // 定义一个channel类型的变量，并为其分配空间（每个channel都是有类型的）

	// 2.管道的长度和容量
	fmt.Println(len(intChan)) // channel的长度：0
	fmt.Println(cap(intChan)) // channel的容量：3

	// 3.管道本身的值
	fmt.Println("intChan:", intChan) // intChan: 0xc000072180，channel是一个引用数据类型，里面存放的是地址值。

	// 4.向管道写入数据
	intChan <- 100

	num := 200
	intChan <- num
	intChan <- 300
	//intChan <- 400  // 注意：向管道写入数据时，不能超过其容量。报错：fatal error: all goroutines are asleep - deadlock!

	// 5.写入数据后的长度和容量
	fmt.Println(len(intChan)) // channel的长度：3
	fmt.Println(cap(intChan)) // channel的容量：3

	// 6.从管道中读取数据
	n1 := <-intChan
	fmt.Println("n1:", n1) // n1: 100（由于管道是一个队列（FIFO），所以先取出来的是最先写入的数据100）

	n2 := <-intChan
	n3 := <-intChan
	//n4 := <-intChan // 读取数据时，也不能超过管道的容量，否则也会报错：fatal error: all goroutines are asleep - deadlock!
	fmt.Println("n2:", n2) // n2: 200
	fmt.Println("n3:", n3) // n3: 300

	//7.如果channel满了，取出数据一个后，就可以继续放入了。（取多少个就可以再放多少个）（类似于水管里面水流的动作）
	intChan <- 1
	intChan <- 2
	intChan <- 3

	// 现在channel已经满了，我们可以取出一个数之后，再往里面放一个数。
	<-intChan     // 取出一个数，不一定要用一个变量来接收，也可以直接丢弃。这里丢弃的就是最先进入管道的1.
	intChan <- 99 // 此时管道里面的内容为：2 3 99
}
