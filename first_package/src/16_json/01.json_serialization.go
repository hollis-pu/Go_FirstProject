package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/**
* Description:
	json序列化的使用
* @Author Hollis
* @Create 2023/9/8 10:40
*/

type Person struct {
	Name  string `json:"name"` // 加上tag之后，序列化为json时，会使用tag中指定的名称而不是原字段名
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	p01 := Person{"Tom", 24, "Tom@gmail.com"}
	fmt.Println("p01:", p01)

	personJson01, err := json.Marshal(&p01) // 将p01序列化为json格式
	if err != nil {
		fmt.Println("json序列化失败！")
		return
	}
	fmt.Println("type of personJson01:", reflect.TypeOf(personJson01)) // []uint8(byte)类型
	fmt.Println("personJson01:", string(personJson01))
}
