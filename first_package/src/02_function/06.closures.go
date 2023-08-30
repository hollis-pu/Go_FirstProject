package main

import "fmt"

/*
*
  - Description:
    闭包的使用
    在这个示例中，`outerFunction` 返回一个闭包函数。
    该闭包函数可以访问并修改 `outerFunction` 的局部变量 `count`。
    每次调用闭包函数，`count` 都会递增并返回新的值，这就形成了一个闭包，可以保留状态。
  - @Author Hollis
  - @Create 2023/8/25 16:16
*/
func countFunc() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func main() {
	counter := countFunc()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}
