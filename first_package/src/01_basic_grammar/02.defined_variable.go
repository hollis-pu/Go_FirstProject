package main

/**
* Description:
	数据类型的定义和查看数据类型
* @Author Hollis
* @Create 2023/8/23 20:41
*/

import (
	"fmt"
	"reflect"
)

func main() {
	var a int32 // 定义变量，若不赋值，则使用默认值0
	a = 10      // 变量赋值

	// 也可以定义时赋值
	var b = 20
	// 也可以不写类型，编译器会自动类型推断
	var c = 20.2
	// 也可以这样定义局部变量，但这种方式不能定义全局变量
	d := 30
	// 变量定义后必须要使用，否则编译不通过
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// 查看变量类型

	// 使用Printf()函数和%T
	fmt.Printf("变量c的类型：%T\n", c) // float64

	//使用reflect包的Typeof()函数
	fmt.Println("变量c的类型：", reflect.TypeOf(c)) // float64
	fmt.Println("变量c的类型：", reflect.TypeOf(d)) // int
}
