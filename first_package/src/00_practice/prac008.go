package main

import "fmt"

/*
*
  - Description:
    题目：已知有个排序好(升序)的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
  - @Author Hollis
  - @Create 2023/8/29 9:50
*/
func main() {
	arr01 := [5]int{1, 3, 4, 5, 6}
	insertData(arr01, 8)
	insertData(arr01, 2)
	insertData(arr01, 3)
	insertData(arr01, -1)

}
func insertData(arr [5]int, data int) {
	var arrTemp [len(arr) + 1]int

	flag := 0 // flag表示是否已经插入数据
	for i := 0; i < len(arr); i++ {
		if flag == 0 { // 未插入时，直接拷贝原数组
			arrTemp[i] = arr[i]
		}
		if arr[i] >= data && flag == 0 { // 第一个大于data的数出现时，data插入到此位置
			arrTemp[i] = data
			flag = 1
		}
		if flag == 0 && i == len(arr)-1 { // 直到最后一个元素都没有插入成功，则插入到最后
			arrTemp[i+1] = data
		}
		if flag == 1 { // 插入元素后，后面的元素向后移动一位
			arrTemp[i+1] = arr[i]
		}
	}
	fmt.Println("arrTemp:", arrTemp)
}
