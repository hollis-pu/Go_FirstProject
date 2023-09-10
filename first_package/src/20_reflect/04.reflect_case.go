package main

import (
	"fmt"
	"reflect"
)

/*
*
  - Description:
    反射的最佳实践：
    使用反射的方式来获取结构体的tag标签，遍历字段的值，修改字段值，调用结构体方法
  - @Author Hollis
  - @Create 2023/9/10 20:33
*/

type Mankind struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func (m *Mankind) Introduce() { // 使用指针接收者的方法，在反射调用时和普通的方法略有不同
	fmt.Printf("My name is %v, %v years old!\n", m.Name, m.Age)
}

func main() {
	m1 := Mankind{"Jerry", 25, "male"}
	callStructure(&m1) // 要修改属性值，这里必须传入地址

}
func callStructure(m *Mankind) { // 指针方式接收形参
	structValue := reflect.ValueOf(m).Elem() // “解引用”

	// 遍历属性和标签
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldName := structValue.Type().Field(i).Name
		tag := structValue.Type().Field(i).Tag.Get("json")

		fmt.Printf("field[%v] is %v, value is %v, tag is %v\n", i, fieldName, field, tag)

		// 设置属性值
		if field.CanSet() {
			if fieldName == "Name" { // 根据字段名修改字段值
				field.SetString("Alice")
			}
			if fieldName == "Age" {
				field.SetInt(28)
			}
			if fieldName == "Gender" {
				field.SetString("female")
			}
		}
	}
	fmt.Println("属性值修改成功！")

	// 调用结构体方法
	// method := structValue.MethodByName("Introduce")  // 这样写无法调用指针接收者方法Introduce
	method := reflect.ValueOf(m).MethodByName("Introduce") // 由于Introduce是指针接收者方法，所以这里应该用指针（解引用之前）来调用该方法。
	if method.IsValid() {
		method.Call(nil)
	} else {
		fmt.Println("方法调用失败！")
	}
}
