package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
*
* Description:
练习2: goroutine+ channel 配合完成排序，并写入文件
要求:
1)开个协程 writeDataToFile，随机生成 1000个数据，存放到文件中。
2)当 writeDataToFile 完成写1000 个数据到文件后,让 sort 协程从文件中读取 1000个数据并完成排序，重新写入到另外一个文件。
3)考察点:协程和管道+文件的综合使用。

4)功能扩展:开10个协程 writeDataToFile，每个协程随机生成 1000 个数据，存放到10文件中。
5) 当10个文件都生成了，让10个 sort 协程从10文件中读取 1000 个数据，并完成排序重新写入到10个结果文件。
* @Author Hollis
* @Create 2023/9/9 11:08
*/

var (
	dataChan  = make(chan string, numCount) // 存放生成的随机数
	boolChan  = make(chan bool, 1)
	writeDone = make(chan bool, 1)
	numCount  = 1000
)

func main() {

	go writeDataToFile()
	go sortData()
	for { // 防止主线程先行退出
		_, ok := <-boolChan
		if !ok {
			break
		}
	}
}
func writeDataToFile() { // 写入数据到文件中
	for i := 0; i < numCount; i++ {
		num := rand.Intn(1001)
		dataChan <- strconv.Itoa(num)
	}
	close(dataChan)
	file, err := os.OpenFile(".\\first_package\\src\\19_channel\\data\\data1000.txt", os.O_CREATE, os.ModePerm)

	if err != nil {
		fmt.Println("文件创建失败！")
		return
	}

	for i := 0; i < numCount; i++ {
		numStr, ok := <-dataChan
		if !ok {
			break
		}
		if i == numCount-1 {
			file.WriteString(numStr)
		} else {
			file.WriteString(numStr + " ")
		}
	}
	fmt.Println("1000个随机数写入完成！")
	defer file.Close()
	writeDone <- true
}

func sortData() {
	<-writeDone
	allFileData, err := os.ReadFile(".\\first_package\\src\\19_channel\\data\\data1000.txt") // 读取文件中的数据
	if err != nil {
		fmt.Println("文件读取失败！")
		return
	}
	numsStr := strings.Split(string(allFileData), " ")
	var nums = make([]int, 0)
	for _, numStr := range numsStr {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}
	sort.Ints(nums) // 将nums进行排序

	sortedFile, _ := os.OpenFile(".\\first_package\\src\\19_channel\\data\\data1000_Sorted.txt", os.O_CREATE, os.ModePerm)
	//defer sortedFile.Close()
	for _, num := range nums {
		_, err := sortedFile.WriteString(strconv.Itoa(num) + " ")
		if err != nil {
			fmt.Println("排序后的数据写入错误！")
			return
		}
	}
	fmt.Println("排序好的数据写入完成！")
	boolChan <- true
	close(boolChan)
}
