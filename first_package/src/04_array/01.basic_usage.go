package main

import "fmt"

/*
*
  - Description:
    数组的基本使用
  - @Author Hollis
  - @Create 2023/8/26 12:59
*/
func main() {
	var arr [5]int32 // 创建长度为5的数组，初始值全为0

	// 遍历数组，方式1：for+len(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i]) // 通过下标访问数组元素
	}

	//给数组元素赋值
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	// 遍历数组，方式2：for+range
	for index, value := range arr {
		fmt.Printf("index=%d value=%d\n", index, value)
	}

	// 数组的地址
	fmt.Printf("arr的地址：%p\n", &arr) // arr的地址：0xc00000e330
	// %p：指针的十六进制表示。
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d]的地址：%v\n", i, &arr[i])
		// arr[0]的地址：0xc00000e330，每次增加4个字节，因为每个元素占4个字节（32位）。
		// arr[1]的地址：0xc00000e334
		// arr[2]的地址：0xc00000e338
	}

	// 几种初始化数组的方式
	// 方式1
	var arr01 [5]int = [5]int{1, 2, 3, 4, 5}
	// 前面的类型可以省略：var arr01 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr01：", arr01)

	// 方式2：类型推断
	// 也可以写成：arr02 := [5]int{1, 2, 3, 4, 5}
	arr02 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr02：", arr02)

	// 方式3：[...]，动态获取初始化数组的长度
	// 这里的[...]是固定写法，不能修改
	arr03 := [...]int{1, 2, 3}   // 这样可以得到长度为3的数组，并给其赋值
	fmt.Println("arr03：", arr03) // arr03： [1 2 3]

	// [...]的方式也可以，根据索引来赋值
	arr04 := [...]int{1: 2, 0: 1, 2: 3}
	fmt.Println("arr04：", arr04) // arr04： [1 2 3]

}
