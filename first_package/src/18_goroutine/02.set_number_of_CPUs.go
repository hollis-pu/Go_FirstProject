package main

import (
	"fmt"
	"runtime"
)

/**
* Description:
	设置程序运行的CPU核心数量
* @Author Hollis
* @Create 2023/9/8 16:19
*/

func main() {
	numCPU := runtime.NumCPU()           // 获取可用的CPU核心数(逻辑上的)
	fmt.Printf("可用的CPU核心数：%d\n", numCPU) // 8

	// 设置程序运行的CPU核心数
	desiredCPU := 4 // 你可以根据需要设置不同的值
	runtime.GOMAXPROCS(desiredCPU)

	// 确认设置后的CPU核心数
	currentCPU := runtime.GOMAXPROCS(0)          // 参数为0时，它会返回当前的CPU核心数，而不会修改设置
	fmt.Printf("当前程序运行的CPU核心数：%d\n", currentCPU) // 4
}
