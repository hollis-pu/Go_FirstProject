package testpackage

import (
	"fmt"
	"testing"
)

/*
*
  - Description:
    单元测试的简单使用：用于单元测试的文件名一定要以“_test”结尾。
  - @Author Hollis
  - @Create 2023/9/8 12:33
*/
func TestCalcSum(t *testing.T) { // 测试用例函数一定要用TestXxx格式，且参数固定为“t *testing.T”
	result := CalcSum(10)
	if result != 55 {
		t.Fatalf("CalcSum(10)的测试用例失败...")
	}
	fmt.Println("CalcSum(10)的测试用例成功...")
}

func TestCalcSub(t *testing.T) {
	result := CalcSub(10, 20)
	if result != -10 {
		t.Fatalf(" CalcSub(10,20)的测试用例失败...")
	}
	fmt.Println(" CalcSub(10,20)的测试用例成功...")
}
