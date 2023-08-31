package main

import "fmt"

/*
*
  - Description:
    快速排序，这个暂时放到后面再学习。（2023.08.31补）
    快速排序的思想：先选定一个比较的基准元素 pivot ，然后遍历数组，依次与pivot进行比较。
    将小于pivot的元素存到 left 切片中，反之则存到 right 切片中。
    然后左右两边的切片都递归调用快速排序函数，直到整个数组都有序。
    需要注意：1.递归的终止条件为切片中的元素个数小于等于1；
    2.在每次递归调用left和right都需要将返回值保存，然后和pivot进行拼接，用以向上返回排序好的切片。
  - @Author Hollis
  - @Create 2023/8/28 21:20
*/
func main() {
	arr01 := [...]int{23, 31, 41, 46, 21, 15, 51, 48, 37, 62} // 共10个数
	arr02 := quickSort(arr01[:])
	fmt.Println("arr02:", arr02)
}
func quickSort(arr []int) []int {
	var left []int
	var right []int
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)-1] // 保存切片中的最后一个元素为pivot
	for _, arrValue := range arr[:len(arr)-1] {
		if arrValue < pivot { // 将pivot作为比较的基准
			left = append(left, arrValue)
		} else {
			right = append(right, arrValue)
		}
	}
	sortedLeft := quickSort(left) // 递归调用quickSort()函数，使sortedLeft是有序的，然后逐层返回
	sortedRight := quickSort(right)
	return append(append(sortedLeft, pivot), sortedRight...) // 将每层得到的sortedLeft、pivot、sortedRight拼接到一起
}
