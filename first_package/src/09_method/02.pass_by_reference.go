package main

import (
	"fmt"
	"math"
)

/*
*
  - Description:
    方法的值传递和引用传递
  - @Author Hollis
  - @Create 2023/8/29 17:47
*/

type Circle struct {
	radius float32
}

func (c Circle) areaCircle01() (area float32) {
	area = math.Pi * c.radius * c.radius
	return
}
func (c *Circle) areaCircle02() (area float32) {
	// 指针接收者：使用指针来接收，方法内部操作的是接收者本身，对对象的修改会导致方法外部对象值的变化
	c.radius = 10
	area = math.Pi * c.radius * c.radius
	return
}

func main() {
	c1 := Circle{3}
	res1 := c1.areaCircle01()
	fmt.Println("res1:", res1) // res1: 28.274334

	c2 := Circle{5}
	res2 := c2.areaCircle02()     // 方法内对radius属性进行了修改
	fmt.Println("c2:", c2.radius) // c2: 10
	fmt.Println("res2:", res2)    // res2: 314.15927
}
