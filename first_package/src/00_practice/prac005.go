package main

import "fmt"

/*
*
  - Description:
    输出100以内的所有素数（只能被1和本身整除的数）
    注：1既不是素数也不是合数
  - @Author Hollis
  - @Create 2023/8/26 11:39
*/
func main() {
	myFunc01()
	myFunc02()
}

// 方式1：使用 continue+标签 的方式
func myFunc01() {
outer:
	for i := 2; i < 100; i++ { // 遍历2-100的所有数字
		for j := 2; j <= i/2; j++ { // 遍历所有可能的因子
			if i%j == 0 {
				continue outer // 直接结束本次外层循环
			}
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// 方式2：使用 flag变量+break 的方式
func myFunc02() {
	for i := 2; i < 100; i++ { // 遍历2-100的所有数字
		isPrime := true
		for j := 2; j <= i/2; j++ { // 遍历所有可能的因子
			if i%j == 0 {
				isPrime = false
				break // 直接结束内层循环
			}
		}
		if isPrime {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
