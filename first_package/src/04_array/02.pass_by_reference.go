package main

import "fmt"

/*
*
  - Description:
    数组的引用传递
  - @Author Hollis
  - @Create 2023/8/26 15:15
*/
func main() {
	arr1 := [3]int{1, 2, 3}
	fmt.Println("arr1=", arr1) // arr1= [1 2 3]
	funcTest(&arr1)
	fmt.Println("arr1=", arr1) // arr1= [10 20 30]
}

func funcTest(arr *[3]int) {
	// 通过引用的方式传递，arr里面存放的是传递过来参数的地址
	// 通过arr加下标的方式可以访问到具体地址的值
	// 这样修改之后，原来数组的值也被修改了
	(*arr)[0] = 10 // (*arr)[0] 和 arr[0] 等价
	arr[1] = 20
	arr[2] = 30
}
