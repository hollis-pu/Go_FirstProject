package first_package

import "fmt"

/**
* Description:
	自定义包下的go代码
* @Author Hollis
* @Create 2023/8/24 15:35
*/

func Myfunc() { // 首字母大写，权限为public，其他包也可以使用
	fmt.Println("这是自定义包下的函数Myfunc！")
}
func myfunc() { // 首字母小写，权限为private，其他包不能使用
	fmt.Println("这是自定义包下的函数myfunc！")
}
