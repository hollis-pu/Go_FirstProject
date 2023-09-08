package main

import (
	"fmt"
	"time"
)

/*
*
  - Description:
    求出小于数字n的所有素数（传统方法：双重循环）（未使用goroutine）
  - @Author Hollis
  - @Create 2023/9/8 13:09
*/
func main() {
	start := time.Now().UnixMilli() // 返回毫秒数
	n := 100000
	result := PrimeNum(n)
	fmt.Println("result", result)
	end := time.Now().UnixMilli()
	fmt.Println("PrimeNum函数执行用时(ms)：", end-start) // 1874ms
}

func PrimeNum(n int) []int {
	var result = make([]int, 0)
	result = append(result, 2)
	for i := 3; i < n; i++ {
		isPrime := true
		if i%2 == 0 {
			continue
		}
		j := 3
		for ; j < i; j += 2 { // j表示i所有可能的因子
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			result = append(result, j)
		}
	}
	return result
}
