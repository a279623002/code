package main

import "fmt"

//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func lengthOfLongestSubstring(s string) int {
	res, left := 0, 0
	tmp := make([]int, 256)
	for right := 0; right < len(s); right++ {
		index := s[right]
		// 如果找到不为0的下标，更新左指针
		left = max(left, tmp[index])
		res = max(res, right - left + 1)
		tmp[index] = right + 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}



func main() {
	fmt.Println(lengthOfLongestSubstring("loddktdji"))
}
