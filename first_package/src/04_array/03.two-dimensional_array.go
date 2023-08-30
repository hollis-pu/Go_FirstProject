package main

import "fmt"

/*
*
  - Description:
    二维数组的基本使用
  - @Author Hollis
  - @Create 2023/8/28 22:14
*/
func main() {
	var arr01 [3][4]int32 // 定义二维数组
	// 二维数组：相当于外层数组里面存的元素又是一维数组（两层数组）

	// 定义时赋值
	arr02 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr02:", arr02) // [[1 2 3] [4 5 6]]

	arr01[1] = [4]int32{1, 1, 1, 1} // arr01[1]的数据类型为[4]int
	fmt.Println("arr01:", arr01)
	fmt.Println(len(arr01)) // 二维数组的长度为外层数组的长度，这里为3

	for i := 0; i < len(arr01); i++ { // 输出外层数组
		fmt.Println(arr01[i])
	}

	for i := 0; i < len(arr01); i++ { // 输出二维数组中的每一个元素
		for j := 0; j < len(arr01[0]); j++ {
			fmt.Printf("%d ", arr01[i][j])
		}
		fmt.Println()
	}

	// 二维数组的地址
	fmt.Printf("arr01[0]的地址：%p\n", &arr01[0]) // 0xc00000e330
	fmt.Printf("arr01[1]的地址：%p\n", &arr01[1]) // 0xc00000e340
	// 0xc00000e330-0xc00000e340=0x10，转换为10进制就是16
	// 由于内层数组的有4个32位（4 byte）的整数，所以arr01[0]和arr01[1]之间相差16 byte
}
