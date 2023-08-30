package main

import "fmt"

/*
*
  - Description:
    指针的简单使用
  - @Author Hollis
  - @Create 2023/8/24 14:43
*/
func main() {
	var num1 int = 23
	// 通过取址符&可以得到变量的地址值
	fmt.Println("num1的地址值为：", &num1) // 0xc00000a0b8

	// var ptr *int = &num1的解析：
	// 1.ptr 是一个指针变量
	// 2.ptr 的类型是 *int
	// 3.ptr 本身的值是 &num1
	//
	var ptr *int = &num1
	fmt.Println("ptr：", ptr) // ptr： 0xc00000a0b8

	// 解引用指针：使用 * 运算符可以获取指针所指向的变量的值。
	fmt.Println("ptr 指向的值：", *ptr) // 23

	// 通过指针可以直接修改其指向的变量的值。
	*ptr = 50
	fmt.Println("*ptr 修改后的值：", *ptr) // 50
}
