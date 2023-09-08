package main

import (
	"fmt"
	"reflect"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/8 20:37
 */

type Person struct {
	name string
	age  int
}

func main() {
	var allChan chan any
	allChan = make(chan any, 10)

	p1 := Person{"Tom", 10}
	p2 := Person{"Jerry", 13}

	allChan <- p1
	allChan <- p2
	allChan <- 10
	allChan <- "jack"

	p1Data := <-allChan
	fmt.Println("p1Data的类型", reflect.TypeOf(p1Data)) // main.Person
	//fmt.Println("p1Data.name:", p1Data.name) // 直接这样使用是会报错的：p1Data.name undefined (type any has no field or method name)

	// 在获取name之前应该使用类型断言，将p1Data转换成Person类型
	person1 := p1Data.(Person)
	fmt.Println("person1.name:", person1.name) // person1.name: Tom

}
