package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
*
  - Description:
    常用的日期时间函数
  - @Author Hollis
  - @Create 2023/8/25 22:34
*/
func main() {

	// 获取当前日期时间
	now := time.Now()
	fmt.Println("当前时间：", now)                   // 2023-08-25 22:36:19.9722907 +0800 CST m=+0.003159601
	fmt.Println("当前年份：", time.Time.Year(now))   // 2023
	fmt.Println("当前月份：", time.Time.Month(now))  // August
	fmt.Println("当前日期：", time.Time.Day(now))    // 25
	fmt.Println("当前小时：", time.Time.Hour(now))   // 22
	fmt.Println("当前分钟：", time.Time.Minute(now)) // 42
	fmt.Println("当前秒数：", time.Time.Second(now)) // 48

	// 日期时间格式化
	currentTime := time.Now()

	// 格式化为 RFC3339 标准时间格式
	fmt.Println("RFC3339:", currentTime.Format(time.RFC3339))
	// 自定义格式化布局
	fmt.Println("Custom format:", currentTime.Format("2006-01-02 15:04:05"))
	// 使用预定义的常用布局
	fmt.Println("Standard layout:", currentTime.Format(time.RFC1123))
	// 格式化为 Unix 时间戳
	unixTimestamp := currentTime.Unix()
	fmt.Println("Unix timestamp:", unixTimestamp)
	// 解析 Unix 时间戳
	unixTime := time.Unix(unixTimestamp, 0)
	fmt.Println("Parsed Unix time:", unixTime.Format("2006-01-02 15:04:05"))

	// 休眠1秒
	time.Sleep(time.Second)

	// 休眠100毫秒
	// 注意，休眠100毫秒只能写成time.Millisecond*100，不能写成time.Second*0.1
	// 因为，Sleep函数能接受的参数类型为Duration类型：Sleep(d Duration)，其中type Duration int64
	//time.Sleep(time.Second * 0.1)  // 报错：0.1 (untyped float constant) truncated to int64
	time.Sleep(time.Millisecond * 100)

	//统计某段程序执行用时
	// Unix()：返回自 UTC 1970 年 1 月 1 日以来经过的秒数
	// 如果需要统计更加精确的时间，可以使用UnixMilli()(毫秒)/UnixMicro()(微秒)/UnixNano()(纳秒)函数，返回的就是对应粒度的时间
	start := time.Now().UnixMilli() // 返回毫秒数
	testFunc()
	end := time.Now().UnixMilli()
	fmt.Printf("testFunc()函数执行用时：%v毫秒\n", end-start) // 1012毫秒

}
func testFunc() {
	str := ""
	for i := 0; i < 10000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
