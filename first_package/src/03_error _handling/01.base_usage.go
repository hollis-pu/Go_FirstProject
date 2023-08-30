package main

import "fmt"

/*
*
  - Description:
    错误处理的基本使用
    错误处理
    Go语言追求简洁优雅，所以，Go语言不支持传统的 try...catch...finally这种处理。
    Go中引入的处理方式为: defer, panic,recover
    这几个异常的使用场景可以这么简单描述: Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理
  - @Author Hollis
  - @Create 2023/8/26 10:32
*/
func main() {
	//test01()
	test02() // 捕获异常之后，代码出错会提示，但不会中断程序，下面的代码继续执行
	fmt.Println("test01()下面的代码继续执行...")
}
func test01() {
	num1 := 10
	num2 := 0
	res := num1 / num2 // 报错：panic: runtime error: integer divide by zero
	fmt.Println("res=", res)
}

func test02() {
	// 使用defer + recover 来捕获异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err) // err: runtime error: integer divide by zero
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2 // 这里报错，本函数后面的代码不再执行
	fmt.Println("res=", res)
}
