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

// 2023.08.28 自己写的快排代码，还是有问题，放在后面在重新学习一下。（2023.08.31已补充）
// 2023.08.31 自己写的这个却递归的中止条件界定得不是很清晰，而且不知道每次递归之后的返回值该如何存放。
// 	      应该让函数有切片类型的返回值，方便递归下层向上层返回。每层得到的返回值也应该用变量进行接收，然后将排序好的left、flag、排序好的right拼接
//            成一个切片，向上返回，这样就能保证每层返回的值都是排序好的了。
func quickSort(arr []int) {
	left := make([]int, 0)
	right := make([]int, 0)

	flag := arr[len(arr)/2]
	for i := 0; i < len(arr); i++ {
		if i == len(arr)/2 {
			continue
		}
		if arr[i] < flag {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	if len(left) > 0 {
		quickSort(left)
	}
	if len(right) > 0 {
		quickSort(right)
	}
}
