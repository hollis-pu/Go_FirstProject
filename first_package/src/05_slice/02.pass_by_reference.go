package main

import "fmt"

/*
*
  - Description:
    切片的引用传递
    切片是引用类型，在函数进行参数传递的时候，传递的是地址
  - @Author Hollis
  - @Create 2023/8/28 15:08
*/
func main() {
	slice01 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice01:", slice01)
	func001(slice01)
	fmt.Println("slice01:", slice01) // slice01: [100 2 3 4 5]
}
func func001(slice []int) {
	slice[0] = 100 // 在函数中对切片的修改，也会影响原切片的值
}
