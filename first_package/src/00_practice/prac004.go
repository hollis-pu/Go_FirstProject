package main

import (
	"fmt"
	"strings"
)

/*
*
  - Description:
    闭包的最佳实践
    1.编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如jpg)，并返回一个闭包。
    2.调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ，则返回 文件名.jpg，
    如果已经有.jpg后缀，则返回原文件名。
    3.要求使用闭包的方式完成
    4.strings.HasSuffix
    func HasSuffix(s, suffix string) bool，判断s是否有后缀字符串suffix。
  - @Author Hollis
  - @Create 2023/8/25 16:45
*/
func main() {
	mySuffix := makeSuffix(".jpg") // 先获得一个闭包

	fmt.Println(mySuffix("hello.jpg")) // hello.jpg
	fmt.Println(mySuffix("winter"))    // winter.jpg
	fmt.Println(mySuffix("pic.png"))   // pic.png.jpg

}
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		} else {
			return name + suffix
		}
	}
}
