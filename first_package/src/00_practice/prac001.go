package main

import "fmt"

/*
*
  - Description:
    求1-100的和
  - @Author Hollis
  - @Create 2023/8/25 9:28
*/
func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1-100的和为：", sum) // 5050
}
