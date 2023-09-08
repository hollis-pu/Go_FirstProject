package testpackage

/**
* Description:
* @Author Hollis
* @Create 2023/9/8 12:25
 */

func CalcSum(n int) int {
	count := 0
	for i := 0; i <= n; i++ {
		count += i
	}
	return count
}

func CalcSub(n1 int, n2 int) int {
	return n1 - n2
}
