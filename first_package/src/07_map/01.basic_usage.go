package main

import "fmt"

/*
*
  - Description:
    map的基本使用
  - @Author Hollis
  - @Create 2023/8/29 11:09
*/
func main() {
	// map使用的几种方式
	// 1.先声明，然后make分配空间
	var map01 map[string]string
	map01 = make(map[string]string, 10)

	// 2.声明，就直接make
	var map02 = make(map[string]string, 10)
	fmt.Println("map02:", map02) // map02: map[]

	// 3.声明，直接赋值（底层会帮我们make分配空间）
	var map03 = map[string]string{"no1": "宋江", "no2": "吴用"}
	fmt.Println("map03:", map03) // map03: map[no1:宋江 no2:吴用]

	// 向map中增加元素
	map01["no1"] = "宋江"
	map01["no2"] = "吴用"
	fmt.Println("map01:", map01) // map[no1:宋江 no2:吴用]

	// 获取map的长度
	fmt.Println("map01的长度", len(map01)) // 2

	// 修改map
	map01["no1"] = "武松"          // map中key不能重复，这里是对 key=no1 的值进行替换
	fmt.Println("map01:", map01) // map[no1:武松 no2:吴用]

	// 删除map中指定的key
	delete(map01, "no2")
	fmt.Println("map01:", map01) // map01: map[no1:武松]

	delete(map01, "no10")        // 当删除的key不存在时，不会进行任何操作，也不会报错
	fmt.Println("map01:", map01) // map01: map[no1:武松]
	// Go语言中不支持将map清空的操作，只能遍历key删除，或者将原map指向新的空map，原map地址的内容将被GC回收。

	// 查找指定的key在map中是否存在
	value01, findRes := map01["no1"] // 第一个返回值：value值，第二个返回值：该key在map中是否存在，存在返回true，否则返回false
	fmt.Println("value01:", value01) // 武松
	fmt.Println("findRes:", findRes) // true

	map01["no3"] = "林冲"
	// 遍历map
	// 使用for+range的方式
	for key, value := range map01 {
		fmt.Printf("key=%s,value=%s", key, value)
		fmt.Println()
	}
}
