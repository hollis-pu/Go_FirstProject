package main

import (
	"fmt"
)

/*
*
  - Description:
    const常量的使用（2023.09.10补）
    常量只能修饰bool、数值类型(int，float系列)、string 类型。
    常量的值必须是在编译时可确定的，不能将变量赋值给常量。
  - @Author Hollis
  - @Create 2023/9/10 16:14
*/
func main() {
	const num = 100
	fmt.Println("num:", num) // 100

	//num = 20 // 常量的值不能被修改，报错：cannot assign to num (neither addressable nor a map index expression)

	// 只能使用基本类型的字面值（如数字、字符串、布尔值）或已声明的常量来初始化常量。
	//num1 := 3
	//const num2 = num1 // 不能将变量赋值给常量，报错：num1 (variable of type int) is not constant

	const (
		a    = iota       // a的值为0
		b                 // b的值为1
		c                 // c的值为2
		d    = "Hello"    // d的值为"Hello"
		e                 // e的值也为"Hello"
		f    = iota       // f的值为5，接着上面的值的个数递增
		g, h = iota, iota // f和g的值都为6。如果在同一行中有多个常量声明，iota只会递增一次
	)
	fmt.Println(a, b, c, d, e, f, g, h) // 0 1 2 Hello Hello 5 6 6
}
