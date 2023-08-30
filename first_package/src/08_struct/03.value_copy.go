package main

import "fmt"

/*
*
  - Description:
    值拷贝的使用
    结构体是值类型，当你将一个结构体变量传递给函数或者赋值给另一个变量时，
    实际上是在进行一次复制操作，复制整个结构体的数据。
  - @Author Hollis
  - @Create 2023/8/29 15:19
*/

type Student struct {
	Name  string
	age   int
	StuNo int
}

func main() {
	stu01 := Student{"Tom", 23, 1001}
	fmt.Printf("stu01:%+v\n", stu01) // stu01:{Name:Tom age:23 StuNo:1001}

	stu02 := stu01 // 值拷贝，修改stu02的值并不会让stu01的值发生变化
	stu02.Name = "Jerry"
	fmt.Printf("stu01:%+v\n", stu01) // stu01:{Name:Tom age:23 StuNo:1001}
	fmt.Printf("stu02:%+v\n", stu02) // stu02:{Name:Jerry age:23 StuNo:1001}

	// 函数参数传递
	func01(stu01)                    // 在函数体里面修改stu01的值，并不会导致函数外面值的变化
	fmt.Printf("stu01:%+v\n", stu01) // stu01:{Name:Tom age:23 StuNo:1001}

	// 值的引用传递（传递地址）
	func02(&stu01)                   // 这样，在函数体里面修改了stu01的值，会导致外面stu01的值也发生改变
	fmt.Printf("stu01:%+v\n", stu01) // stu01:{Name:Zoe age:23 StuNo:1001}

}
func func01(stu Student) {
	stu.Name = "Zoe"
}
func func02(stu *Student) {
	stu.Name = "Zoe"
}
