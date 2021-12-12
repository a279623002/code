package main

import "fmt"

//输入：x = -121
//输出：false
//解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
func isPalindrome(x int) bool {
	// 小于0 || 11110
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}
	a := 0
	for x > a {
		a = a * 10 + x % 10
		x /= 10
	}
	return x == a || x == a / 10
}

func main() {
	fmt.Println(isPalindrome(332233))
}
