package main

import "fmt"

/*
*
  - Description:
    map的值传递
  - @Author Hollis
  - @Create 2023/8/29 13:31
*/
func main() {
	map01 := map[string]int{
		"key1": 100,
		"key2": 200,
	}
	fmt.Println("map01:", map01) // map01: map[key1:100 key2:200]
	funcMap(map01)
	fmt.Println("map01:", map01) // map01: map[key1:999 key2:200]
	// map是引用类型，在函数体中修改map，也会导致原map值的变化
}
func funcMap(mapData map[string]int) {
	mapData["key1"] = 999
}
