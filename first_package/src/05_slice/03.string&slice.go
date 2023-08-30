package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
*
  - Description:
    string与切片的结合
  - @Author Hollis
  - @Create 2023/8/28 15:15
*/
func main() {

	// 对string进行切片
	s1 := "hello,world!"
	slice01 := s1[0:5] // 这样得到的slice01还是string类型，而不是切片类型
	//必须使用[]byte()函数来对截取得到的字符串进行显式转换
	slice02 := []byte(s1[0:5])
	fmt.Println("slice01的类型:", reflect.TypeOf(slice01)) // string
	fmt.Println("slice02的类型:", reflect.TypeOf(slice02)) // []uint8

	// string是不可修改的
	//s1[0] = 'a'  // 报错：cannot assign to s1[0]

	// 如果需要修改字符串，可以先将string -> []byte /或者 rune -> 修改 -> 重写转成string。
	s2 := "hello,world!"
	slice03 := []byte(s2) // 这里需要将s2转换成切片类型
	// 注意，使用s2[:]得到的还是string类型，必须通过显示转换才可以
	slice03[0] = 'a'
	s2 = string(slice03)   // 再将切片类型转换回string类型
	fmt.Println("s2:", s2) // s2: aello,world!

	// 将含有存有多个字符串的切片按指定分隔符连接
	// 使用strings.Join()函数
	slice04 := []string{"hello", "world"}
	s3 := strings.Join(slice04, " ")
	fmt.Println("s3=", s3) // s3= hello world
}
