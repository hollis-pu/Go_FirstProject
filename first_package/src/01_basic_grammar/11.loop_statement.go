package main

import "fmt"

/*
*
  - Description:
    循环语句的使用
  - @Author Hollis
  - @Create 2023/8/24 19:30
*/
func main() {
	// 1. 基本for循环的使用
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
	}

	// 2. for + range 的使用
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 3. 无限循环的使用
	//for {
	//	fmt.Println("hello")
	//}

	// 4. 类似与while的条件循环：在循环体里面写变量的增加语句，循环条件写在for后面
	num := 0
	for num < 5 {
		fmt.Println("num=", num)
		num++
	}
}
