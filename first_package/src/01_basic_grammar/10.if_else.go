package main

import "fmt"

/*
*
  - Description:
    if...else...的使用
  - @Author Hollis
  - @Create 2023/8/24 18:59
*/
func main() {
	score := 200
	if score < 60 {
		println("不及格！")
	} else if score < 70 { // 注意：这里的else只能写在上一行，不能换行写
		println("合格！")
	} else if score < 90 {
		println("良好！")
	} else if score <= 100 {
		println("优秀！")
	} else {
		println("成绩异常！")
	}

	// 注意，下面的程序只会输出一个 ok1。
	// 因为在使用多个条件分支时，因为只会执行第一个满足条件的分支。
	var n int = 10
	if n > 9 {
		fmt.Println("ok1")
	} else if n > 6 {
		fmt.Println("ok2")
	} else if n > 3 {
		fmt.Println("ok3")
	} else {
		fmt.Println("ok4")
	}
}
