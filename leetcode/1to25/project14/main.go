package main

import "fmt"

//输入：strs = ["flower","flow","flight"]
//输出："fl"
func longestCommonPrefix(strs []string) string {
	i := 0
	for i < len(strs[0]) {
		match := false
		for _, v := range strs[1:] {
			if i >= len(v) || strs[0][i] != v[i] {
				match = false
				break
			}
			match = true
		}
		if !match {
			break
		}
		i++
	}
	return strs[0][0:i]
}

func main() {
	strs := []string{

	}
	fmt.Println(longestCommonPrefix(strs))
}