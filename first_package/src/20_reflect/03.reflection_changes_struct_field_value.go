package main

import (
	"fmt"
	"reflect"
)

/**
* Description:
	反射修改结构体变量的值：必须传入地址，否则报错：panic: reflect: reflect.Value.Set using unaddressable value
* @Author Hollis
* @Create 2023/9/10 20:24
*/

type Student struct {
	Name string
	Age  int
}

func main() {
	Student := Student{"Alice", 30}
	valueType := reflect.ValueOf(&Student).Elem() // 注意：这里必须传入地址（然后通过Elem()获取到该地址存入的内容），才能修改结构体字段的值。

	if valueType.Kind() == reflect.Struct {
		for i := 0; i < valueType.NumField(); i++ {
			field := valueType.Field(i)
			fieldType := field.Type()
			fieldName := valueType.Type().Field(i).Name

			fmt.Printf("Field Name: %s\n", fieldName)
			fmt.Printf("Field Type: %s\n", fieldType)

			// 检查字段是否可被设置
			if field.CanSet() {
				fmt.Printf("Field Value (before): %v\n", field.Interface())

				// 修改字段的值
				if fieldType == reflect.TypeOf(0) {
					newValue := reflect.ValueOf(40)
					field.Set(newValue)
					fmt.Printf("Field Value (after): %v\n", field.Interface())
				} else if fieldType == reflect.TypeOf("") {
					newValue := reflect.ValueOf("Bob")
					field.Set(newValue)
					fmt.Printf("Field Value (after): %v\n", field.Interface())
				}
			} else {
				fmt.Printf("Field is not settable\n")
			}
			fmt.Println()
		}
	}
}
