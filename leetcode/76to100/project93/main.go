package main

import (
	"fmt"
	"strconv"
)

//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
func restoreIpAddresses(s string) []string {
	// 用小数点正确分割
	// 1. 只能有三个小数点
	// 2. 分组只有4个元素
	// 3. 元素不超过3个
	// 4. 元素不大于255
	// 5. 元素0开头只能强制分割，不能01这种
	res := []string{}

	var fn func(int)
	fn = func(i int) {
		if i == len(s) {

		} else {

		}
	}
	fn(0)
	return res
}

func check(s string) bool {
	i, _ := strconv.Atoi(s)
	if i > 255 {
		return true
	}
	return false
}

func main() {
	fmt.Println(restoreIpAddresses("010010"))
}
