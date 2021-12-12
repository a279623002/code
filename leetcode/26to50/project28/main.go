package main

import "fmt"

//输入：haystack = "hello", needle = "ll"
//输出：2
func strStr(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}
	res := -1
	for first := 0; first < len(haystack) - len(needle) + 1; first++ {
		if haystack[first:first+len(needle)] == needle {
			return first
		}
	}
	return res
}

func main() {
	fmt.Println(strStr("a", "a"))
}
