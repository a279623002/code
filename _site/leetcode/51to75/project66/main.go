package main

import "fmt"

//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
func plusOne(digits []int) []int {
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		carry, digits[i] = (digits[i] + 1)/10, (digits[i] + 1)%10
		if carry == 0 {
			break
		}
	}
	if carry > 0 {
		return append([]int{carry}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}
