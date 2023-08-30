package main

import (
	"fmt"
	"reflect"
)

/**
* Description:
	字符和字符串的使用
		和其他编程语言不同的是，Golang底层存储字符串是用字节数组存的，而不是字符数组。

		字符串使用UTF-8编码来存储Unicode字符。
		在UTF-8编码中，每个字符由一个或多个字节组成。
		ASCII字符（0-127）仍然使用单个字节表示，因此与ASCII兼容。
		非ASCII字符（如汉字）使用多个字节表示，具体取决于字符的Unicode代码点。

		对于单个字符，如果其码值小于等于255（在Ascii表里面），则可以直接保存到byte
		如果其对应的码值大于255，则byte存放不下，可以考虑用int类型保存
* @Author Hollis
* @Create 2023/8/24 9:52
*/

func main() {

	c1 := 'h'
	fmt.Println(c1)                           // 这里会输出h的ascii码104，而不是字符'h'
	fmt.Println("c1的类型：", reflect.TypeOf(c1)) // 系统推断的类型为：int32，也可以显式地指定类型为byte
	fmt.Printf("c1 = %c\n", c1)               // 使用%c来格式化输出，c1=h

	var c2 byte = 'h'
	fmt.Println(c2)                           // 这里会输出h的ascii码104，而不是字符'h'
	fmt.Println("c2的类型：", reflect.TypeOf(c2)) // uint8（等价于byte）
	fmt.Printf("c2 = %c\n", c2)               // h

	// 存储中文的情况
	c3 := '中'
	fmt.Println(c3)                           // 输出Unicode码值：20013
	fmt.Println("c3的类型：", reflect.TypeOf(c3)) // int32
	fmt.Printf("c3 = %c\n", c3)               // 使用%c来格式化输出，c2=中

	// var c4 byte = '中'
	// 报错：cannot use '中' (untyped rune constant 20013) as byte value in variable declaration (overflows)
	// 因为byte类型存不下Unicode码值为20013的”中“字符，可以考虑用int来存。

	var str1 string = "hello"
	//str1[0] = 'a' // 报错：cannot assign to str1[0] (neither addressable nor a map index expression)
	fmt.Printf("%s\n", str1)
	fmt.Printf("%d\n", len(str1))
	fmt.Printf("%c\n", str1[0])
	fmt.Printf("%c\n", str1[1])

	// 当字符串的拼接需要换行时，必须把“+”号写在上一行而不能写在下一行
	// 因为编译器在每行的末尾自动加上“;”，而如果我们写了“+”号，就不会再添加“;”。
	str2 := "hello" + "hello" +
		"hello" + "hello"
	fmt.Println(str2)

}
