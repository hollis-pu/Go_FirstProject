package main

import (
	"fmt"
	"math"
)

/**
* Description:
	接口的基本使用
	Go语言中，接口的实现是隐式的，只要类型的方法集合包含了接口中的所有方法，就被视为实现了该接口。
* @Author Hollis
* @Create 2023/8/30 13:49
*/

// Shape 接口定义了图形的绘制方法
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体代表了矩形
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle 结构体代表了圆形
type Circle struct {
	Radius float64
}

// Area 隐式实现Shape接口（即：实现Shape的所有方法）
// Area 计算矩形的面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形的周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area 隐式实现Shape接口
// Area 计算圆形的面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆形的周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建一个包含不同图形的切片
	shapes := []Shape{
		Rectangle{Width: 4, Height: 5},
		Circle{Radius: 3},
	}

	// 遍历图形并打印其面积和周长
	for _, shape := range shapes {
		fmt.Printf("Shape: %T\n", shape)
		fmt.Printf("Area: %.2f\n", shape.Area()) // 在具体的调用时，体现多态性
		fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
		fmt.Println("-------------")
	}
}
