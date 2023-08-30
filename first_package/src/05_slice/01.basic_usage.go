package main

import "fmt"

/*
*
  - Description:
    切片的基本使用
  - @Author Hollis
  - @Create 2023/8/28 13:53
*/
func main() {
	// 切片的创建，并初始化
	// 切片的长度是可以变化的，因此切片是一个可以动态变化数组。
	slice1 := []int{10, 20, 30}
	fmt.Println("slice1=", slice1) // slice1= [10 20 30]

	// 向切片中追加数据（一个或多个）
	// append之后，生成新的切片，原来的切片值不变
	slice2 := append(slice1, 100, 200, 300)
	fmt.Println("slice1=", slice1) // slice1= [10 20 30]
	fmt.Println("slice2=", slice2) // slice2= [10 20 30 100 200 300]

	// 向切片中追加切片
	slice3 := []int{9, 99, 999}
	slice4 := append(slice1, slice3...) // ... 表示对切片的解压缩
	fmt.Println("slice4=", slice4)      // slice4= [10 20 30 9 99 999]

	// 使用 make() 函数创建一个切片，长度为 2，容量为 10
	// make的方式会在底层创建一个数组，该数组对外不可见，即只能通过 slice去访间各个元素
	slice5 := make([]int, 2, 10)
	slice5[0] = 100
	fmt.Println("slice5=", slice5) // slice5= [100 0]

	arr01 := [...]int{0, 1, 2, 3, 4, 5}

	// 对数组的切片，就是对数组指定位置的引用
	slice6 := arr01[1:4] // [1 2 3]，不包含结束位置
	fmt.Println("slice6=", slice6)

	fmt.Println("arr01[1]的地址：", &arr01[1])   //  0xc000122038
	fmt.Println("slice6[0]的地址：", &slice6[0]) //  0xc000122038
	// 这二者是相等的，说明切片是对数组的引用

	// 修改切片的值，也会导致原数组的值被修改
	slice6[0] = 100
	fmt.Println("arr01=", arr01)   // arr01= [0 100 2 3 4 5]
	fmt.Println("slice6=", slice6) // slice01= [100 2 3]

	// 复制切片
	// 如果是直接赋值，则是将原切片结构中存的全部信息拷贝给新的切片
	// 修改任意一个切片中的值，都会导致另一个切片值的变化（两个切片只是别名的关系）
	slice7 := []int{1, 2, 3, 4, 5}
	slice8 := slice7
	slice8[0] = 100
	fmt.Println("slice7=", slice7) // slice7= [100 2 3 4 5]
	fmt.Println("slice8=", slice8) // slice8= [100 2 3 4 5]

	// 可以使用copy()函数来是实现切片值的拷贝
	// 这样得到的两个切片就是相互独立的，不会彼此影响
	slice9 := []int{1, 2, 3, 4, 5}
	slice10 := make([]int, len(slice9))
	copy(slice10, slice9)
	slice10[0] = 100
	fmt.Println("slice9=", slice9)   // slice9= [1 2 3 4 5]
	fmt.Println("slice10=", slice10) // slice10= [100 2 3 4 5]

	slice11 := make([]int, 2, 3)
	slice12 := append(slice11, 1, 2, 3, 4, 5)
	fmt.Println("slice12:", slice12)
}
