package main

import (
	"fmt"
	"math"
)

/*
*
  - Description:
    跳水比赛，8个评委打分。运动员的成绩是 8 个成绩取掉一个最高分，去掉一个最低分，剩下的6个分数的平均分就是最后得分。
    使用一维数组实现如下功能:
    (1) 请把打最高分的评委和最低分的评委找出来。
    (2) 找出最佳评委和最差评委。最佳评委就是打分和最后得分最接近的评委。最差评委就是打分和最后得分相差最大的。
  - @Author Hollis
  - @Create 2023/8/29 10:27
*/
func main() {
	arr01 := [8]int{34, 67, 83, 72, 61, 85, 76, 90}

	maxIndex := 0 // 最大值的索引
	minIndex := 0 // 最小值的索引
	var sum = 0.0
	for i := 0; i < len(arr01); i++ {
		if arr01[i] > arr01[maxIndex] {
			maxIndex = i
		}
		if arr01[i] < arr01[minIndex] {
			minIndex = i
		}
		sum += float64(arr01[i])
	}
	// 去掉最高分和最低分，求平均值
	average := (sum - float64(arr01[maxIndex]+arr01[minIndex])) / float64(len(arr01)-2)
	fmt.Println("最高分的评委：", maxIndex)
	fmt.Println("最低分的评委：", minIndex)
	fmt.Println("平均分：", average)

	worstIndex := 0 // 最差评委的索引
	bestIndex := 0  // 最佳评委的索引
	// 遍历数组，找出差值最大的和最小的
	for i := 0; i < len(arr01); i++ {
		if math.Abs(float64(arr01[i])-average) > math.Abs(float64(arr01[worstIndex])-average) {
			worstIndex = i
		}
		if math.Abs(float64(arr01[i])-average) < math.Abs(float64(arr01[bestIndex])-average) {
			bestIndex = i
		}
	}

	fmt.Println("最佳评委：", bestIndex)
	fmt.Println("最差评委：", worstIndex)
}
