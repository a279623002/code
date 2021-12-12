package main

import "fmt"

//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	dict := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	res := []string{}
	// 		a 			b 			c
	// d	e	f 	d	e	f 	d	e	f
	// def def def def def def def def def
	var fn func(int, string)
	fn = func(i int, s string) {
		if i == len(digits) {
			res = append(res, s)
		} else {
			index := string(digits[i])
			for _, v := range dict[index] {
				fn(i+1, s+string(v))
			}
		}
	}
	fn(0, "")

	return res
}

func main() {
	fmt.Println(letterCombinations(""))
}
