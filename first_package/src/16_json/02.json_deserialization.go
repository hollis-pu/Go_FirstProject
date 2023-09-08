package main

import (
	"encoding/json"
	"fmt"
)

/*
*
  - Description:
    json反序列化
  - @Author Hollis
  - @Create 2023/9/8 10:40
*/

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	jsonData := []byte(`{"name":"John","age":30,"email":"john@example.com"}`) // json.Unmarshal()函数接收的是byte类型的切片，所以这里需要用[]byte()来进行类型转换。

	var stu Student
	err := json.Unmarshal(jsonData, &stu) // 第一个参数的类型为：[]byte
	if err != nil {
		fmt.Println("json反序列化失败，", err)
		return
	}
	fmt.Println("stu:", stu) // stu: {John 30 john@example.com}
}
