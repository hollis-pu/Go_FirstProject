package main

import (
	"fmt"
	"os"
)

/*
*
  - Description:
    写入内容到文件
  - @Author Hollis
  - @Create 2023/9/7 20:18
*/
func main() {
	var file01 *os.File
	_, err := os.Stat("test_write_content.txt") // 获取文件的状态，如果文件不存在，则返回错误信息。
	if err != nil {                             // 若文件不存在，则创建该文件
		fmt.Println("文件不存在，现创建此文件！")
		file01, err = os.Create("test_write_content.txt")
		if err != nil {
			fmt.Println("文件创建失败！")
			return
		}
		fmt.Println("文件创建成功！")
	}
	// 开始写入文本内容到文件
	fmt.Println("开始写入文件...")
	_, err = file01.WriteString("hello world!")
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
	fmt.Println("写入文件成功！")

}
