package main

import "fmt"

/*
*
  - Description:
    函数的值传递机制
    基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值
    如果参数是指针类型，那么传递的是指针的副本，但仍然是值传递。
    如果在函数内部通过指针修改了指向的数据，那么函数外部的数据也会受到影响，因为它们指向同一块内存。
  - @Author Hollis
  - @Create 2023/8/25 12:12
*/
func main() {
	num := 10
	myFunc01(num)
	fmt.Println("myFunc01 num =", num) // num=10，调用函数myFunc01()并不会改变num的值

	myFunc02(&num)                     // 传入num的地址
	fmt.Println("myFunc02 num =", num) // num=20，调用函数myFunc02()会改变num的值

}

// 基本数据类型的值传递，不会改变外部变量的值
func myFunc01(num int) {
	num = 20
}

// 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&。
// 函数内以指针的方式操作变量。从效果上看类似引用。
func myFunc02(num *int) { // num接收传递过来的地址，通过指针来访问该地址的值
	*num = 20
}
