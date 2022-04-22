package main

import "fmt"

//输入：x = 8
//输出：2
//解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
func mySqrt(x int) int {
	res := 0
	left, right := 0, x
	for left <= right {
		m := (left+right)/2
		if m * m > x {
			right = m-1
		} else {
			left = m+1
			res = m
		}
	}
	return res
}

func main() {
	fmt.Println(mySqrt(9))
}
