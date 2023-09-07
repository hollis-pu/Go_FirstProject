package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

/*
*
  - Description:
    文件的基本操作：创建文件
  - @Author Hollis
  - @Create 2023/9/7 19:48
*/
func main() {
	// 创建文件
	file01, err := os.Create("go_create_file.txt") // 创建一个空的文件，放在当前工作目录下。

	// 可以通过`os.Getwd()`函数获取当前工作目录，一般为当前项目的目录，而不是当前包的目录。
	workPath, _ := os.Getwd()
	fmt.Println(workPath) // E:\code\golang\Go_FirstProject

	if err != nil {
		fmt.Println("创建文件失败：", nil)
		return
	}
	defer file01.Close()
	fmt.Println("Type of file:", reflect.TypeOf(file01)) // *os.File

	filePath, err := filepath.Abs(file01.Name()) // 获取文件的绝对路径（需传入一个string类型的文件名，而不是file类型的文件）
	fmt.Println(filePath)                        // E:\code\golang\Go_FirstProject\go_create_file.txt
}
