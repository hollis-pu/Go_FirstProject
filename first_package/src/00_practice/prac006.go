package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
*
  - Description:
    数组的应用案例
    1) 创建一个byte类型的26个元素的数组，分别放置A-Z。
    使用for循环访问所有元素并打印出来。提示: 字符数据运算'A'+1->B
    2) 请求出一个数组的最大值，并得到对应的下标。
    3) 请求出一个数组的和和平均值。for-range
    4) 随机生成五个数构成一个数组，并将其反转。
  - @Author Hollis
  - @Create 2023/8/28 12:35
*/
func main() {
	func001()

	arr1 := [...]int{16, 3, 9, 13, 32, 61, 81, 19, 42}
	func002(arr1)
	func003(arr1)

	func004()
}

// 创建一个byte类型的26个元素的数组，分别放置A-Z。
func func001() {
	var bytes [26]byte
	bytes[0] = 'A'
	for i := 0; i < 26; i++ {
		bytes[i] = byte('A' + i)
		fmt.Printf("%c ", bytes[i])
	}
	fmt.Println()
}

// 请求出一个数组的最大值，并得到对应的下标。
func func002(arr [9]int) (arrayMax int, indexMax int) {
	arrayMax = arr[0]
	indexMax = 0
	for i := 0; i < len(arr); i++ {
		if arrayMax < arr[i] {
			arrayMax = arr[i]
			indexMax = i
		}
	}
	fmt.Println("数组的最大值为：", arrayMax)
	fmt.Println("数组最大值的下标为：", indexMax)
	return
}

// 请求出一个数组的和和平均值。for-range
func func003(arr [9]int) (arrayAver float32) {
	var count float32 = 0.0
	for i := 0; i < len(arr); i++ {
		count += float32(arr[i])
	}
	arrayAver = count / float32(len(arr))
	fmt.Println("数组的平均值为：", arrayAver)
	return
}

// 随机生成五个数构成一个数组，并将其反转。
func func004() {
	// 生成伪随机数
	// rand.Intn(n)：返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
	// 这样得到的随机数一直是不变的，如果想要变化，需要使用rand.Seed(time.Now().UnixNano())指定随机数种子
	// rand.NewSource()是原rand.Seed()函数的替代函数
	var arr [5]int
	rand.NewSource(time.Now().UnixNano())
	// 得到5个随机数，并存入数组中
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("原数组：", arr)

	// 将原数组反转
	for i := 0; i < len(arr)/2; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = temp
	}
	fmt.Println("反转后：", arr)
}
