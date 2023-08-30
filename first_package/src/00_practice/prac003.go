package main

import "fmt"

/**
* Description:
	请编写一个函数swap(n1 int,n2 int) 可以交换n1和n2的值（如果函数签名是这样，那么无法进行交换）
	只能通过指针的方式才能交换：swap(n1 *int,n2 *int)
* @Author Hollis
* @Create 2023/8/25 14:16
*/

func main() {
	n1 := 20
	n2 := 30
	swap01(n1, n2)
	fmt.Println("n1 n2=", n1, n2)

	swap02(&n1, &n2) // 传入地址，然后通过指针访问指定地址的值并进行修改
	fmt.Println("n1 n2=", n1, n2)
}
func swap01(n1 int, n2 int) { // 这种写法是无法交换n1和n2的值的
	temp := n1
	n1 = n2
	n2 = temp
}

func swap02(n1 *int, n2 *int) { // 这种写才能交换n1和n2的值
	temp := *n1
	*n1 = *n2
	*n2 = temp
}
