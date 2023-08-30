package main

import (
	"Go_FirstProject/first_package/src/10_factory_pattern/student"
	"fmt"
)

/**
* Description:
	工厂模式的使用
* @Author Hollis
* @Create 2023/8/29 20:48
*/

func main() {

	// 使用student的工厂函数，来创建其对象
	s1 := student.Factory("Tom", 23, "男")
	fmt.Println("s1:", *s1) // s1: {Tom 23 男}

	// 获取对象的属性值
	fmt.Println("name=", s1.GetName(), ",age=", s1.GetAge(), ",gender=", s1.GetGender())
}
