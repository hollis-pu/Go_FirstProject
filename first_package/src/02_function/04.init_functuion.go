package main

import "fmt"

/*
*
  - Description:
    init函数的使用：
    `init()` 函数在程序启动时自动被调用，无需手动调用。
    每个导入的包中的 `init()` 函数都会被执行。

	如果一个文件同时包含全局变量定义，init() 函数和 main() 函数，则执行的流程：
	全局变量定义 -> init() 函数 -> main() 函数
  - @Author Hollis
  - @Create 2023/8/25 14:26
*/
// 如果有全局变量，则编译器先进行全局变量的定义
var a = test01()

func test01() int {
	fmt.Println("test()函数的调用...")
	return 90
}

func main() {
	fmt.Println("main()函数的调用...")
}

// init()函数在main()函数之前调用
func init() { // 在main()函数之前调用
	fmt.Println("init()函数的调用...")
}

//输出结果为：
//test()函数的调用...
//init()函数的调用...
//main()函数的调用...
