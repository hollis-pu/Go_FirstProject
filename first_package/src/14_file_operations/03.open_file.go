package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
*
  - Description:
    打开文件
    Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；
    对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。
  - @Author Hollis
  - @Create 2023/9/7 20:48
*/
func main() {
	file01, err := os.Open("open_file_case.txt")
	if err != nil {
		fmt.Println("文件打开失败：", err)
	}
	defer file01.Close()

	// 1. 带缓冲的
	// 创建一个 *Reader，是带缓冲的，可以读一部分，处理一部分。这样就可以处理一些比较大的文件。
	//const (
	//	defaultBufSize = 4096
	//)
	reader := bufio.NewReader(file01)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //  io.EOF表示文件的末尾，注意当读取到最后一行时，就会返回err，但是最后一行的数据需要再下面输出。
			fmt.Print(str)
			break
		}
		fmt.Print(str)
	}
	fmt.Println()
	fmt.Println("文件读取结束！")

	// 2. 一次性读取全部内容
	content, err := os.ReadFile("open_file_case.txt")
	fmt.Println(string(content))

}
