package main

import (
	"fmt"
	"reflect"
)

/*
*
  - Description:
    反射获取结构体的属性和方法
  - @Author Hollis
  - @Create 2023/9/10 19:24
*/

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p Person) SayHello() {
	fmt.Println("hello,my name is", p.Name)
}
func (p Person) CalSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	p1 := Person{"Tom", 25, "male"}

	value := reflect.ValueOf(p1)
	fmt.Println(value.Type()) // main.Person
	fmt.Println(value.Kind()) // struct

	if value.Kind() == reflect.Struct {
		// 遍历结构体的字段
		for i := 0; i < value.NumField(); i++ { // 通过 value.NumField() 可以获取到结构体的属性数量
			field := value.Field(i) // value.Field(i): 获取结构体第i个属性的值
			fieldType := field.Type()
			fieldName := value.Type().Field(i).Name // 返回结构体属性的名称
			// 注意：reflect.Value的Field()方法返回的是Value类型。
			// 而reflect.Type的Field()方法返回的是StructField类型，StructField中存放了所反射的结构体的属性名、标签等信息。
			fmt.Printf("Field Name: %s, Field Type: %s, Field Value: %v\n", fieldName, fieldType, field.Interface())

			// field.Set(reflect.ValueOf(100))  // 要想修改结构体变量的属性值，在传入必须传递变量的地址，否则报错：
			// panic: reflect: reflect.Value.Set using unaddressable value
			// 详见03代码

		}
		// 调用结构体的无参方法
		method := value.MethodByName("SayHello")       // 通过函数名称调用结构体中的函数
		if method.IsValid() && method.CanInterface() { // 如果方法合法
			methodValue := method.Call(nil) // 调用方法并传递接收者，call(nil)表示不传入任何参数
			fmt.Println(methodValue)        // 输出方法的返回值（如果有的话）
		} else {
			fmt.Println("Method not found or invalid")
		}

		//调用结构体的带参方法
		method1 := value.MethodByName("CalSum")
		var params []reflect.Value                   // 声明一个[]reflect.Value的切片，来存放函数的参数列表
		params = append(params, reflect.ValueOf(10)) // 指定参数1的值
		params = append(params, reflect.ValueOf(20)) // 指定参数2的值
		methodValue1 := method1.Call(params)         // 调用方法并传入参数列表，将返回值列表存入methodValue1中
		fmt.Println(methodValue1[0].Interface())     // 查看第一个返回值的结果
	}
}
