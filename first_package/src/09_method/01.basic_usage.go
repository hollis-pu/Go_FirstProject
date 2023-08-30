package main

import "fmt"

/*
*
  - Description:
    方法的基本使用
  - @Author Hollis
  - @Create 2023/8/29 17:24
*/

type Student struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	s1 := Student{"Tom", 23, "男"}
	s1.learnMethod() // Tom 正在学习...

	s2 := Student{"Rose", 24, "女"}
	s2.getGender() // Rose 的性别为：女

	//num1 := 100
	//num1.learnMethod() // 报错，learnMethod() 方法只能被 Student 类型的变量来调用
}
func (s Student) learnMethod() {
	fmt.Printf("%s 正在学习...\n", s.Name)
}
func (s Student) getGender() {
	fmt.Printf("%s 的性别为：%s\n", s.Name, s.Gender)
}
