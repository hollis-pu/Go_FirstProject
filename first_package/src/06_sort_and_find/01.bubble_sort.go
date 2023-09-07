package main

import "fmt"

/*
*
  - Description:
    冒泡排序
  - @Author Hollis
  - @Create 2023/8/28 16:53
*/
func main() {
	arr01 := [...]int{23, 31, 41, 46, 21, 15, 51, 48, 37, 62} // 共10个数
	bubbleSort(&arr01)
	fmt.Println("bubbleSort arr01:", arr01)
	//quickSort(arr01[:])
	//fmt.Println("quickSort arr01:", arr01)
}

// 冒泡排序：每轮使1个数在正确的位置上，需要双层循环
// 外层循环：控制循环的轮数，每轮排好1个数
// 内层循环：从外层循环的i值开始，表示第i个元素是即将需要参与排序的数
func bubbleSort(arr *[10]int) {
	for i := 0; i < len(arr)-1; i++ { // 控制循环的轮数
		for j := i; j > 0; j-- { // // 控制实际参与排序的数字
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
}

// 2023.08.28 自己写的快排代码，还是有问题，放在后面在重新学习一下。
//func quickSort(arr []int) {
//	left := make([]int, 0)
//	right := make([]int, 0)
//
//	flag := arr[len(arr)/2]
//	for i := 0; i < len(arr); i++ {
//		if i == len(arr)/2 {
//			continue
//		}
//		if arr[i] < flag {
//			left = append(left, arr[i])
//		} else {
//			right = append(right, arr[i])
//		}
//	}
//	if len(left) > 0 {
//		quickSort(left)
//	}
//	if len(right) > 0 {
//		quickSort(right)
//	}
//}
