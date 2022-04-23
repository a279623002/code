package main

import (
	"fmt"
)

//输入：s = "Hello World "
//输出：5
func lengthOfLastWord(s string) int {
	res, n := 0, len(s)-1
	for i := n; i >= 0; i-- {
		// 一开始遇到空格不加个数
		if s[i] == ' ' {
			// 有个数后再遇到空格就返回
			if res != 0 {
				break
			}
		} else {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(lengthOfLastWord("hello world "))
}
