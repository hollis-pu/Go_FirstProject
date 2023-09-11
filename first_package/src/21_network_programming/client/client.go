package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

/*
*
  - Description:
    socket编程，客户端的程序
    实现了一个简单的TCP客户端，可以与服务器建立连接并向服务器发送文本数据。
    客户端不断等待用户输入，直到用户输入 "exit" 时退出循环，关闭连接，并退出程序。
    这样，客户端可以与服务器进行交互性通信。
  - @Author Hollis
  - @Create 2023/9/11 13:26
*/
func main() {
	fmt.Println("客户端开始连接...")
	host := "localhost"
	port := 8888

	// 与服务器端建立连接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	defer conn.Close()
	fmt.Printf("conn 成功=%v time=%s\n", conn, time.Now().Format("2006-01-02 15:04:05"))

	reader := bufio.NewReader(os.Stdin)

	// 客户端可以一直向服务器端写入数据，每次写入一行。当输入exit时，关闭连接。
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}

		if line == "exit\r\n" { // 注意，这里要加上\r\n
			fmt.Println("客户端退出！")
			break
		}
		n, err := conn.Write([]byte(line)) // 通过连接向服务器端写入数据
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		fmt.Printf("客户端发送了 %d 字节的数据 %s\n", n, time.Now().Format("2006-01-02 15:04:05"))
	}
}
