package main

import (
	"fmt"
	"strconv"
)

//输入: a = "1010", b = "1011"
//输出: "10101"
func addBinary(a string, b string) string {
	res, carry := "", 0
	al, bl := len(a), len(b)
	for i, j := al-1, bl-1; i >=0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		m, n := 0, 0
		if i >= 0 {
			m = int(a[i] - '0')
		}
		if j >= 0 {
			n = int(b[j] - '0')
		}
		sum := m+n+carry
		sum, carry = sum%2, sum/2
		res = strconv.Itoa(sum) + res
	}

	return res
}

func main() {
	fmt.Println(addBinary("1010", "1011"))
}
