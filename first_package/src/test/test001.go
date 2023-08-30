package main

import (
	"encoding/json"
	"fmt"
)

/**
* Description:
* @Author Hollis
* @Create 2023/8/29 11:39
 */
// 课堂练习: 演示一个key-value 的value是map的案例
// 比如:我们要存放3个学生信息,每个学生有 name和sex 信息
func main() {
	map01 := map[string]map[string]string{
		"stu01": {"name": "Joe", "sex": "female"},
		"stu02": {"name": "George", "sex": "male"},
		"stu03": {"name": "Tom", "sex": "male"},
	}
	fmt.Println("map01:", map01)
	jsonStr, _ := json.Marshal(map01)
	fmt.Println("jsonStr:", string(jsonStr))

}
