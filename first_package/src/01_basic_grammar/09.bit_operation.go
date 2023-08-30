package main

import "fmt"

/*
*
  - Description:
    位移运算
  - @Author Hollis
  - @Create 2023/8/24 18:02
*/
func main() {
	//位移运算后，a,b,c,d 结果是多少
	var a int = 1 >> 2  // 0
	var b int = -1 >> 2 // -1
	var c int = 1 << 2  // 4
	var d int = -1 << 2 // -4
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	var i1 int = 0b1101
	fmt.Println("i1=", i1)
}
