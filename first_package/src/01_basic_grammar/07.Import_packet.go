package main

import (
	"Go_FirstProject/first_package/src/first_package"
)

/*
*
  - Description:
    引入自定义包的使用
  - @Author Hollis
  - @Create 2023/8/24 15:31
*/
func main() {
	// 调用自定义包下面的函数：包名.函数名(...)
	//first_package.myfunc() // 报错：无法在当前软件包中使用未导出的 函数 'myfunc'
	first_package.Myfunc()
}
