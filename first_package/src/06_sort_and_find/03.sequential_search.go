package main

import "fmt"

/*
*
  - Description:
    顺序查找
  - @Author Hollis
  - @Create 2023/8/28 21:20
*/
func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	sequentialSearch(arr, 22)
}
func sequentialSearch(arr [8]int, findElem int) {
	index := -1
	for i := 0; i < len(arr); i++ {
		if findElem == arr[i] {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("没有查找到元素：", findElem)
	} else {
		fmt.Printf("元素 %d 的下标为：%d", findElem, index)
	}
}
