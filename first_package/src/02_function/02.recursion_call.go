package main

import "fmt"

/*
*
  - Description:
    函数的递归调用
    1. 用递归实现斐波那契数列
    2. 关于递归理解的一个案例
  - @Author Hollis
  - @Create 2023/8/25 10:21
*/
func main() {
	for i := 0; i < 20; i++ {
		fmt.Printf("f(%d) = %d\n", i, fibonacci(i))
	}

	test(4)
}

// 递归调用fibonacci()函数，实现斐波那契数列
func fibonacci(i int) int {
	if i <= 1 {
		return 1
	} else {
		return fibonacci(i-1) + fibonacci(i-2)
	}
}

// 帮助理解的递归的案例
// 思考：下面的代码输出结果是什么？
func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

// 解析：
//
//1.第1次调用：test(4)
//
//   n=4，n>2，进入 if 语句块：
//
//   ​	n--：n=3
//
//   ​	test(3)：调用test(3)
//
//2.第2次调用：test(3)
//
//   n=3，n>2，进入 if 语句块：
//
//   ​	n--：n=2
//
//   ​	test(2)：调用test(2)
//
//3.第3次调用：test(2)
//
//   n=2，n>2，不满足 if 条件，往下执行输出语句：
//
//   ​	fmt.Println("n=", n)：输出n=2
//
// **【注意】到这里，整个递归调用的过程并没有结束，而是回到前一个递归层次（这是很容易出错的地方）。**
//
//4.回到第2次调用，继续执行下面的输出语句：
//
//   ​	fmt.Println("n=", n)：输出n=2（注意，这里的局部变量n值为2）
//
//   再回到前一个递归层次。
//
//5.回到第1次调用，继续执行下面的输出语句：
//
//   ​	fmt.Println("n=", n)：输出n=3
//
//   这已经是第1次调用了，整个递归调用结束。
//
//综上，代码的输出结果为：
//
//```
//n= 2
//n= 2
//n= 4
//```
//
//注意，虽然 `test` 函数在递归时修改了 `n` 的值，但每个递归调用都有自己的局部变量副本，因此修改不会影响到其他调用。
