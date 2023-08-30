package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

/*
*
  - Description:
    接口的最佳实践
    实现对Student结构体切片的排序: sort.Sort(data Interface)
    要使用sort.Sort(data Interface)方法，对Student的指定字段进行排序
    Sort()方法需要传入一个接口类型（或实现了该接口的类型（多态））
    Interface是Go官方定义的一个接口，里面有三个方法：
    type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
    }
    现在要实现对Student类型的排序，也就是要在Sort()函数中传入Student类型的参数，
    那么就需要让Student结构体实现Interface接口。（即实现上面的三个方法）
  - @Author Hollis
  - @Create 2023/8/30 15:03
*/

type Student struct {
	name  string
	age   int
	score int
}

// StuSlice 定义一个切片类型，存放多个学生
type StuSlice []Student

// 实现Interface接口中的三个方法：Len、Less、Swap

// Len 切片中存放的学生数量
func (s StuSlice) Len() int {
	return len(s)
}

// Less 指定排序规则（按年龄的升序排序）
func (s StuSlice) Less(i, j int) bool {
	return s[i].age < s[j].age
}

// Swap 将切片中的元素进行交换
func (s StuSlice) Swap(i, j int) {
	//temp := s[i]
	//s[i] = s[j]
	//s[j] = temp
	// 上面的3条语句等价于：
	s[i], s[j] = s[j], s[i]
}
func main() {
	s1 := StuSlice{}

	for i := 0; i < 10; i++ {
		s1 = append(s1, Student{"stu" + strconv.Itoa(i),
			int(float32(rand.Intn(101))*0.1 + 15), // 生成15-25之间的随机数
			rand.Intn(101)})                       // 生成0-100之间的随机数
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", s1[i])
	}
	fmt.Println("-----排序后-----")
	sort.Sort(s1)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", s1[i])
	}
}
