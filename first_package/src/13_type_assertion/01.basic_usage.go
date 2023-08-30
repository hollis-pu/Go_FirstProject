package main

import "fmt"

/**
* Description:
	类型断言的使用
	断言可以检查接口值的实际类型，并根据需要进行类型转换。
* @Author Hollis
* @Create 2023/8/30 19:00
*/

type Person struct {
	Name string
	Age  int
}
type Car struct {
	Brand string
	Price int
}

func main() {
	var i01 interface{}
	p01 := Person{"Tom", 23}
	i01 = p01 // 多态

	// 类型断言
	i02, isOk := i01.(Person)                      // 断言返回的一个返回值为转换之后的值，第二个返回值为断言是否成功。
	fmt.Printf("i02的值为：%v，i02的类型为：%T\n", i02, i02) // i02的值为：{Tom 23}，i02的类型为：main.Person
	fmt.Println("info:", isOk)                     // true

	//i03 := i01.(Car) // 当断言失败时，如果只接收一个返回值，会抛出一个panic。
	//// panic: interface conversion: interface {} is main.Person, not main.Car

	i03, isOk := i01.(Car) // 当断言失败时，当接收两个返回值时不会抛出panic，但第二个返回值为 false。
	// 我们可以通过这种方式来防止程序抛出panic。通过判断isOk的值是否为true，来确实是否断言成功。
	fmt.Printf("i03的值为：%v，i03的类型为：%T\n", i03, i03) // i03的值为：{ 0}，i03的类型为：main.Car
	fmt.Println("info:", isOk)                     // false

	// 还可以使用 switch 进行类型分支
	// 这样可以根据实际的类型，在case中进行不同的操作（比如调用各自类型的方法等等）
	switch v := i01.(type) {
	case Car:
		fmt.Printf("i01的类型为：%T，值为：%v\n", v, v)
	case Person:
		fmt.Printf("i01的类型为：%T，值为：%v\n", v, v) // i01的类型为：main.Person，值为：{Tom 23}
	case int:
		fmt.Printf("i01的类型为：%T，值为：%v\n", v, v)
	default:
		fmt.Printf("i01的类型未知！\n")
	}
}
