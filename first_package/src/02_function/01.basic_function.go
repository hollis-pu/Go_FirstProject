package main

import "fmt"

/*
*
  - Description:
    函数的基本使用
  - @Author Hollis
  - @Create 2023/8/25 9:35
*/
func main() {
	a, b := 10, 20
	fmt.Println(add1(a, b)) // 30
	fmt.Println(add2(a, b)) // 30

	// 多参数函数的调用，使用多个变量来接收返回值
	sum, sub := addSub(a, b)
	fmt.Println("sum is:", sum)
	fmt.Println("sub is:", sub)

	// 若不想接收函数的某个返回值，也可以使用匿名变量“_”（不能所有的返回值都使用匿名变量接收）
	sum1, _ := addSub(a, b)
	fmt.Println("sum1 is:", sum1) // 30

	sum2, _ := addSub1(a, b)
	fmt.Println("sum2 is:", sum2) // 30

	// 可变参数函数的调用
	fmt.Println(funcSum(1, 2, 3))       // 6
	fmt.Println(funcSum(1, 2, 3, 4, 5)) // 15
}

// 定义一个函数，包含两个参数，后面是返回值名称及其类型，用()括起来。
func add1(a int, b int) (sum int) {
	return a + b
}

// 如果参数类型相同，则可以省略，只写后者的类型，返回值的名称也可以省略不写
func add2(a, b int) int {
	return a + b
}

// Go语言的函数支持多个返回值
func addSub(a, b int) (int, int) {
	return a + b, a - b
}

// Go语言支持返回值命名
func addSub1(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

// 可变参数列表，使用“...”来表示可变参数列表
func funcSum(numbers ...int) int { // 求任意数量的参数之和
	sum := 0
	for _, number := range numbers { // 使用for...range...遍历参数列表
		sum += number
	}
	return sum
}
