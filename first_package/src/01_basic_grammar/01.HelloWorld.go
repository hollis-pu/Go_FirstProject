package main

import "fmt" // 导入 fmt 包

func main() { // 程序的入口main函数，注意，这个大括号不能换行写
	// 输出hello world
	fmt.Println("hello world!")

	// 转义字符\n表示换行
	fmt.Println("你好 \n world!")

	// 使用反引号`，不解析转义字符
	fmt.Println(`你好 \n world!`) // 你好 \n world!

}
