package main

import "fmt"

/**
  - Description:
    继承的基本使用
  - @Author Hollis
  - @Create 2023/8/30 11:24

*/

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) ShowInfo() {
	fmt.Printf("学生姓名=%v 年龄=%v 成绩=%v\n", s.Name, s.Age, s.Score)
}

func (s *Student) SetScore(score int) {
	s.Score = score
}

type Pupil struct {
	Student // 嵌入Student的匿名结构体，实现继承的功能
}

type Graduate struct {
	Student // 继承Student
}

// Learning 大学生特有的方法

func (g *Graduate) Learning() {
	fmt.Println("大学生正在学习")
}

func main() {
	pupil01 := &Pupil{}
	//pupil01.Student.Name = "Tom"  // 可以简写为：pupil01.Name = "Tom"
	pupil01.Name = "Tom"
	pupil01.Age = 13
	pupil01.SetScore(90)
	pupil01.ShowInfo()

	graduate01 := &Graduate{}
	graduate01.Name = "Tomas"
	graduate01.Age = 25
	graduate01.Learning() // 调用大学生特有的方法Learning()
	graduate01.SetScore(80)
	graduate01.ShowInfo()
}
