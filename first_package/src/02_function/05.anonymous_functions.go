package main

import "fmt"

/*
*
  - Description:
    匿名函数的使用
    两种方式：
    1.定义时直接调用。
    2.将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数。
  - @Author Hollis
  - @Create 2023/8/25 16:08
*/
func main() {
	// 方式1
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(20, 30)
	fmt.Println("result:", result) // 50

	// 方式2：此时的resultFunc为一个函数变量，后面可以通过resultFunc来直接调用函数
	resultFunc := func(num1 int, num2 int) int {
		return num1 + num2
	}
	fmt.Println("resultFunc:", resultFunc(40, 50)) // 90

}
