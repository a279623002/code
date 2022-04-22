package main

import "fmt"

//输入: s = "III"
//输出: 3
func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for k, _ := range s {
		if k < len(s) - 1 && dict[s[k]] < dict[s[k + 1]] {
			res -= dict[s[k]]
		} else {
			res += dict[s[k]]
		}
	}

	return res
}

func main() {
	fmt.Println(romanToInt("IV"))
}