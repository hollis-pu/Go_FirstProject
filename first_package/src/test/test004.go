package main

import (
	"fmt"
	"strconv"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/3 21:23
 */
func main() {
	var num int
	fmt.Scanln(&num)
	count, factorSlice := test(num)
	fmt.Println(count)
	var resultStr string
	for i := 0; i < len(factorSlice); i++ {
		if i < len(factorSlice)-1 {
			resultStr += strconv.Itoa(factorSlice[i]) + "*"
		} else {
			resultStr += strconv.Itoa(factorSlice[i])
		}
	}
	fmt.Println(resultStr)
}
func test(num int) (int, []int) {
	var factor []int
	for i := 2; i*i < num; i++ {
		if num%i == 0 {
			factor = append(factor, i)
			//fmt.Printf("%d ", i)
		}
	}
	count := 1
	maxLength := 1
	for i := 0; i < len(factor)-1; i++ {
		if factor[i+1] == factor[i]+1 {
			count++
		} else {
			if count > maxLength {
				maxLength = count
			}
			count = 1
		}
	}
	var resultSlice []int
	for i := 0; i < len(factor); i++ {
		if factor[i]+maxLength-1 == factor[i+maxLength-1] {
			resultSlice = factor[i : i+maxLength]
			break
		}
	}
	return maxLength, resultSlice
}
