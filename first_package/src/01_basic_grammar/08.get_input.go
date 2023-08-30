package main

import "fmt"

/*
*
  - Description:
    通过键盘获取输入的使用：fmt.Scan()、fmt.Scanln() 和 fmt.Scanf()
  - @Author Hollis
  - @Create 2023/8/24 17:43
*/
func main() {
	var name string
	var age int

	// 1. fmt.Scan()
	// 用于按空格分隔读取多个值，将用户输入的数据按空格分隔开，并依次存储到提供的变量中。
	// 直到输入enter结束。
	// 用户输入的数据的数量必须与提供的变量数量相匹配。
	//fmt.Print("请输入姓名和年龄：")
	//fmt.Scan(&name, &age)
	//
	//fmt.Println("姓名为：", name)
	//fmt.Println("年龄为：", age)

	// 2. fmt.Scanln()
	// 它会一直读取，直到遇到换行符（Enter键）为止，然后将输入的数据存储到提供的变量中。
	//fmt.Print("请输入姓名：")
	//fmt.Scanln(&name)
	//fmt.Print("请输入年龄：")
	//fmt.Scanln(&age)
	//
	//fmt.Println("姓名为：", name)
	//fmt.Println("年龄为：", age)

	// 3. fmt.Scanf()
	// 类似于 C 语言的 scanf()，它允许你指定输入的格式，然后根据格式读取用户输入的数据。
	fmt.Print("请输入姓名和年龄：")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println("姓名为：", name)
	fmt.Println("年龄为：", age)
}
