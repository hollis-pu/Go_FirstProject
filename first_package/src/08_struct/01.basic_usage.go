package main

import "fmt"

/**
* Description:
	结构体的基本使用
* @Author Hollis
* @Create 2023/8/29 14:36
*/

type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	// 结构体变量的使用
	// 方式1：先声明，后赋值
	var cat1 Cat                   // 创建一个cat类型的变量，其属性使用默认值
	fmt.Printf("cat1:%+v\n", cat1) // cat1:{Name: Age:0 Color:}

	// 属性赋值
	cat1.Name = "小花"
	cat1.Age = 2
	cat1.Color = "橘色"
	fmt.Printf("cat1:%+v\n", cat1) // cat1:{Name:小花 Age:2 Color:橘色}

	// 方式2：定义时直接赋值
	cat2 := Cat{"小乖", 3, "黑色"}
	fmt.Printf("cat2:%+v\n", cat2) // cat2:{Name:小乖 Age:3 Color:黑色}

	//方式3-new
	//案例: var cat *cat3 = new (Cat)
	var cat3 *Cat = new(Cat)
	//因为 cat3 是一个指针，因此标准的给字段赋值方式
	//(*cat3).Name ="花花"，也可以这样写 cat3.Name ="花花"
	//原因: go的设计者 为了程序员使用方便，底层会对 cat3.Name ="花花" 进行处理
	//会给 cat3 加上 取值运算 (*cat3).Name ="花花"
	(*cat3).Name = "花花" // 标准写法
	cat3.Name = "小白"    // 简便写法

	(*cat3).Age = 4
	cat3.Age = 5
	fmt.Printf("cat3:%+v\n", cat3) // cat3:&{Name:小白 Age:5 Color:}

	// 方式4：*和&
	var cat4 *Cat = &Cat{"大橘", 3, "橘色"}
	fmt.Printf("cat4:%+v\n", cat4) // cat4:&{Name:大橘 Age:3 Color:橘色}
}
