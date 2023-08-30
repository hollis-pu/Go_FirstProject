package main

import (
	"errors"
	"fmt"
)

/*
*
  - Description:
    手动抛出自定义错误
    Go程序中，也支持自定义错误， 使用 errors.New 和 panic 内置函数。
    1) errors.New("错误说明")，会返回一个 error 类型的值，表示一个错误。
    2) panic()内置函数 接收一个 interface{} 类型的值(也就是任何值了)作为参数。
    可以接收 error 类型的变量，输出错误信息，并退出程序。
  - @Author Hollis
  - @Create 2023/8/26 11:03
*/
func main() {
	testMyErr()
	fmt.Println("main()中代码继续执行...")
}

func readConf(name string) (err error) {
	if name != "config.ini" {
		return errors.New("读取配置文件出错！") // 当文件名不匹配时，返回一个 error 类型的值
	}
	return nil
}

func testMyErr() {
	err := readConf("config.ini")
	if err != nil {
		panic(err) // 使用panic()函数，手动抛出一个panic，会被系统捕获并中止程序
	} else {
		fmt.Println("配置文件读取成功！")
	}
	fmt.Println("testMyErr()继续执行...")
}
