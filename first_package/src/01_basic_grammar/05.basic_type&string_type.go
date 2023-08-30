package main

import (
	"fmt"
	"strconv"
)

/**
* Description:
* @Author Hollis
* @Create 2023/8/24 13:48
 */
func main() {
	// 基本数据类型转换为string
	// 注意不能直接使用string(num)的方式，这样会把整数值num转换为对应的 Unicode 字符
	var g1 int = 65
	var g2 string = string(g1) // 注意，这种方式是将65转换为其对应的unicode字符A

	fmt.Println("g1=", g1) // 65
	fmt.Println("g2=", g2) // A

	// 要将65转换为string类型，需要使用下面的方式
	// 方式1：fPrintf()函数
	var g1Str1 string = fmt.Sprintf("%d", g1)
	fmt.Println("g1Str1=", g1Str1) // g1Str1= 65

	// 也可以转换其他类型为string，如浮点数，布尔类型的值等
	// 浮点数转成string
	ff := 12.23
	ffStr := fmt.Sprintf("%f", ff)
	fmt.Println("ffStr=", ffStr) // ffStr=12.230000

	// bool转成string
	bb := true
	bbStr := fmt.Sprintf("%t", bb)
	fmt.Println("bbStr=", bbStr) // bbStr=true

	// 方式2：使用 strconv.Itoa()（整数到字符串）
	// Itoa()等价于FormatInt(int64(i), 10)
	g1Str2 := strconv.Itoa(g1)
	fmt.Println("g1Str2=", g1Str2) // g1Str2= 65

	// 浮点数转成string
	ffStr2 := strconv.FormatFloat(ff, 'f', 3, 64)
	// 第一个参数：要转换的浮点数
	// 第二个参数：要转换的格式
	// 第三个参数：保留的小数点位数
	// 第四个参数：浮点数的位数
	fmt.Println("ffStr2=", ffStr2) // ffStr2= 12.230

	// string转成基本数据类型，使用strconv.ParseXxx()函数
	// strconv.ParseXxx()函数有两个返回值：(i int64, err error)
	// 我们现在只需要第一个返回值，那么第二个返回值可以使用“_”接收，表示忽略该返回值。

	// string转成int，可以使用Atoi()函数，等价于ParseInt(s, 10, 0)
	str1 := "12345"
	str1Int1, _ := strconv.Atoi(str1)
	fmt.Println("str1Int1=", str1Int1) // str1Int1= 12345

	// 使用ParseInt()函数
	str1Int2, _ := strconv.ParseInt(str1, 10, 64)
	fmt.Println("str1Int2=", str1Int2) // str1Int2= 12345

	// string转成浮点数：使用ParseFloat()函数
	str2 := "23.34"
	str2Float, error1 := strconv.ParseFloat(str2, 64)
	fmt.Println("str2Float=", str2Float) // str2Float= 23.34
	fmt.Println("error1=", error1)       // error1= <nil>

	// 如果string不能转换成基本数据类型，则第一个参数得到的值为0，第二个参数的得到的是报错的信息
	str3 := "hello"
	str3Int, err := strconv.ParseInt(str3, 10, 64)
	fmt.Println("str3Int=", str3Int) // “hello”不能转换成Int，str3Int值为0
	fmt.Println("err=", err)         // err= strconv.ParseInt: parsing "hello": invalid syntax

}
