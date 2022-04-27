package main

import (
	"fmt"
	"strconv"
)

func fn(str1, str2 string) string {
	n1, n2 := len(str1)-1, len(str2)-1
	res, carry := "", 0
	for n1 >= 0 || n2 >= 0 || carry > 0 {
		r1, r2 := 0, 0
		if n1 >= 0 {
			r1 = int(str1[n1]-'0')
			n1--
		}
		if n2 >= 0 {
			r2 = int(str2[n2]-'0')
			n2--
		}
		sum := r1 + r2 + carry
		res = strconv.Itoa(sum%2) + res
		carry = sum/2
	}
	return res
}

func main() {
	fmt.Println(fn("1111", "1011"))
}
