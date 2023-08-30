package main

import "fmt"

/*
*
  - Description:
    String()方法的使用
  - @Author Hollis
  - @Create 2023/8/29 18:49
*/

type Person struct {
	name   string
	age    int
	gender string
}

func (p Person) String() string {
	return fmt.Sprintf("Person[name:%s, age:%d, gender:%s]", p.name, p.age, p.gender)
}

func main() {
	p1 := Person{"Tom", 23, "男"}
	// 实现String()方法后，在输出Person对象时，会自动调用String()方法，按照其中定义的格式输出
	fmt.Println("p1:", p1) // p1: Person[name:Tom, age:23, gender:男]
}
