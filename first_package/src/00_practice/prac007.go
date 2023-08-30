package main

import "fmt"

/*
*
  - Description:
    切片的课堂练习题
    说明: 编写一个函数 fbn(n int)，要求完成
    1.可以接收一个n int
    2.能够将斐波那契的数列放到切片中
    3.提示，斐波那契的数列形式：arr[0] = 1; arr[1] = 1; arr[2]=2; arr[3]= 3; arr[4]=5; arr[5]=8
  - @Author Hollis
  - @Create 2023/8/28 16:38
*/
func main() {
	slice01 := fbn(10)
	fmt.Println("slice01:", slice01)
}
func fbn(n int) (arr []int) {
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		if i <= 1 {
			arr[i] = 1
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return
}
