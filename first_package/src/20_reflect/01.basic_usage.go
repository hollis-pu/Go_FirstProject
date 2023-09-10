package main

import (
	"fmt"
	"reflect"
)

/*
*
  - Description:
    反射的基本使用
  - @Author Hollis
  - @Create 2023/9/10 17:30
*/
func main() {
	// 1.获取类型信息
	// 使用reflect.TypeOf()函数可以获取一个值的类型信息
	num1 := 100
	t1 := reflect.TypeOf(num1)
	fmt.Println("t1:", t1) // t1:int

	// 也可以通过reflect.Value类型的Type()方法来获取变量的类型
	t2 := reflect.ValueOf(num1).Type()
	fmt.Println("t2:", t2) // t2:int

	// 2.获取值信息
	// 使用reflect.ValueOf()函数可以获取一个值的反射对象Value
	v1 := reflect.ValueOf(num1)
	fmt.Printf("v1:%v\n", v1)      // v1:100
	fmt.Printf("v1 type:%T\n", v1) // v1 type:reflect.Value。注意：v1的类型为reflect.Value，而不是int。

	// v2 := v1 + 10  // 不能直接将v1和整数相加，需要v1取出里面的int类型的值。如下：
	v2 := v1.Int() + 10
	fmt.Println("v2:", v2) // v2: 110

	// 3.改变变量的值
	// 通过reflect.Value对象的Set()方法，可以修改变量的值
	// v1.SetInt(200) // 报错：panic: reflect: reflect.Value.SetInt using unaddressable value
	// 只有当变量是可寻址的（addressable）时，才能修改其值。
	v11 := reflect.ValueOf(&num1) // 传值时传入地址
	v12 := v11.Elem()             // 然后通过Elem()取出其地址的值（类似于指针的解引用）
	v12.SetInt(200)               // 使用SetXxx()修改变量的值
	fmt.Println("v1:", v12)       // v12: 200
	fmt.Println("num1:", num1)    // num1: 200。可以看到原变量num1的值也发生了改变

	// 4.获取种类（kind）信息
	// Type和Kind的区别：Type反映类型，Kind反映种类，前者是后者的子集。具体如下：
	// Type是reflect.Value对象的一个方法，用于获取值的类型信息。通过Type方法，你可以知道一个值的确切类型，例如，是一个int32、string、Person、mySlice等等。它提供了完整的类型信息，包括名称和包路径。
	// Kind是reflect.Value对象的一个方法，用于获取值的底层基础类型的分类。它返回一个reflect.Kind类型的枚举值，表示值的底层类型分类，如int、string、struct、slice等。它用于分类基本类型，但不提供详细的类型信息。
	k1 := v1.Kind()
	fmt.Println("k1:", k1) // k1: int

	// 5.类型检查
	if v1.Kind() == reflect.Int { // 这里调用的就是reflect.Kind类型的枚举值
		fmt.Println("值的类型为int")
	}

}
