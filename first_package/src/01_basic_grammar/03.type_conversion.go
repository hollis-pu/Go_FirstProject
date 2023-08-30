package main

import (
	"fmt"
	"reflect"
)

/*
*
  - Description:
    数据类型的转化：Go语言只能进行显式的类型转换，不支持隐式类型转换
  - @Author Hollis
  - @Create 2023/8/23 21:18

*/

func main() {
	var a int = 10
	//a = 10.2  // 注意，这样写会报错，因为a为int类型，不能保存float类型的数据
	fmt.Println("a=", a)

	var b float32 = 21.2
	b = 20 // 这里不会报错（因为20可以看做20.0），但b的类型还是float，而不是int
	fmt.Println("b=", b)
	fmt.Println("TypeOf(b)=", reflect.TypeOf(b)) // TypeOf(b)= float32

	c := 23
	d := 2
	e := c / d                                   // 这里e的类型，系统会推断为int而不是float，因为参与运算的c和d都是int类型
	fmt.Println("e=", e)                         // e= 11
	fmt.Println("TypeOf(e)=", reflect.TypeOf(e)) // TypeOf(e)= int

	// go语言中也不允许int类型和float类型相除
	//x := 100
	//y := 33.0
	//result := x / y // 报错：invalid operation: x / y (mismatched types float64 and int)
	//fmt.Println("result=", result)

	// 可以使用int()，float32()等函数，来进行强制类型转换
	// 但应该注意的是，由于Go语言中不存在隐式类型转换，所以强制类型转换不一定能得到我们想要的结果。
	var f_e float32 = float32(c / d) // 这里得到的f_e值为11，而不是11.5。
	// 因为这里是将c/d得到的11转换成float类型，而不是11.5。

	fmt.Println("f_e=", f_e)                         // 11
	fmt.Println("TypeOf(f_e)=", reflect.TypeOf(f_e)) // TypeOf(f_e)= float32

	// 注意：被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
	var t1 int = 20
	var t2 float32 = float32(t1)
	fmt.Println("t1的类型：", reflect.TypeOf(t1)) // int
	fmt.Println("t2的类型：", reflect.TypeOf(t2)) // float32

}
