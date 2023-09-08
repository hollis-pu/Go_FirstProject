package main

import (
	"fmt"
	"os"
)

/*
*
  - Description:
    通过os.Args获取命令行的参数
  - @Author Hollis
  - @Create 2023/9/7 22:05
*/
func main() {
	// 使用os.Args来获取命令行参数，os.Args[0]是程序的名称
	fmt.Println("程序名称:", os.Args[0])

	// os.Args[1:] 包含所有的命令行参数，不包括程序名称
	fmt.Println("命令行参数:", os.Args[1:])

	// 可以通过索引访问特定的命令行参数
	if len(os.Args) > 1 {
		fmt.Println("第一个命令行参数:", os.Args[1])
	}
}
