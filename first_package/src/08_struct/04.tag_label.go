package main

import (
	"encoding/json"
	"fmt"
)

/**
* Description:
	tag标签的使用
* @Author Hollis
* @Create 2023/8/29 16:41
*/

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	p1 := Person{"Tom", 23, "male"}
	fmt.Printf("p1:%+v\n", p1) // p1:{Name:Tom Age:23 Gender:male}
	// json.Marshal() 中使用到了反射
	jsonStr, err := json.Marshal(p1) // 这样得到的第一个返回值为byte数组，第二个返回值为错误信息
	if err != nil {
		fmt.Println("json处理出错：", err)
	}
	// 需要将byte数组转换成string
	// 添加tag标签之前，得到的json键的首字母为大写，添加tag之后，首字母为小写
	fmt.Println("jsonStr:", string(jsonStr)) // jsonStr: {"name":"Tom","age":23,"gender":"male"}
}
