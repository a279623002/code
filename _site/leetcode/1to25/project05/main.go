package main

import "fmt"

//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end, left, right := 0, 0, 0, 0
	for i := 0; i < len(s); i++ {
		left, right = expand(i, i, s)
		if right - left > end - start {
			start, end = left, right
		}
		left, right = expand(i, i+1, s)
		if right - left > end - start {
			start, end = left, right
		}
	}
	return s[start:end+1]
}

func expand(left, right int, s string) (int, int) {
	for ;left >= 0 && right < len(s) && s[left] == s[right];  {
		left--
		right++
	}
	return left+1, right-1
}

func main() {
	fmt.Println(longestPalindrome(""))
}
