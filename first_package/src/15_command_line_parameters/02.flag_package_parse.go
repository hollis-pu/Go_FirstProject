package main

import (
	"flag"
	"fmt"
)

/*
*
  - Description:
    使用flag包来解析命令行参数
  - @Author Hollis
  - @Create 2023/9/7 22:13
*/
func main() {
	// 声明命令行参数
	var (
		name    string
		age     int
		verbose bool
	)

	// 使用flag包创建命令行参数，传递变量的指针以便将值存储在变量中
	flag.StringVar(&name, "name", "Guest", "设置名称")
	flag.IntVar(&age, "age", 0, "设置年龄")
	flag.BoolVar(&verbose, "verbose", false, "启用详细输出")

	// 解析命令行参数
	flag.Parse()

	// 打印解析后的参数值
	fmt.Println("名称:", name)
	fmt.Println("年龄:", age)
	fmt.Println("详细输出:", verbose)
}
