package main

import "fmt"

/**
* Description:
	猴子吃桃子问题
	有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个!
	以后每天猴子都吃其中的一半，然后再多吃一个。
	当到第十天时，想再吃时 (还没吃)，发现只有1个桃子了。
	问题:最初共多少个桃子？
* @Author Hollis
* @Create 2023/8/25 11:14
*/

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(f(i)) // 最初的数量：i=10时，值为1534
	}
}

func f(day int) int {
	// f(day int) 表示的是从今天开始为第1天，往回数的第day天，猴子还剩多少个桃子
	if day == 1 {
		return 1
	} else {
		return 2 * (f(day-1) + 1) // 根据题意，倒推出前一天数量和后一天数量的表达式
	}
}
