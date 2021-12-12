package main

import "fmt"

//输入: num = 3
//输出: "III"
func intToRoman(num int) string {
	dict := []struct {
		index int
		str   string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	res := ""
	for _, v := range dict {
		for num >= v.index {
			num -= v.index
			res += v.str
		}
	}
	return res
}

func main() {
	fmt.Println(intToRoman(0))
}
