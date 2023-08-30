package main

import "fmt"

/*
*
  - Description:
    二分查找
  - @Author Hollis
  - @Create 2023/8/28 21:28
*/
func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	binarySearch(&arr, 0, len(arr)-1, 8)
}
func binarySearch(arr *[8]int, leftIndex int, rightIndex int, findElem int) {

	// 当leftIndex大于rightIndex时，表示未找到指定元素
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	middleIndex := (leftIndex + rightIndex) / 2
	middleElem := arr[middleIndex]

	if findElem < middleElem {
		binarySearch(arr, leftIndex, middleIndex-1, findElem)
	} else if findElem > middleElem {
		binarySearch(arr, middleIndex+1, rightIndex, findElem)
	} else {
		fmt.Println("找到了，下标为：", middleIndex)
	}
}
