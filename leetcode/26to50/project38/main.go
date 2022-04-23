package main

import (
	"fmt"
	"strconv"
)

//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//第一项是数字 1
//描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
//描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
//描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
//描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	frontStr := countAndSay(n-1)
	sum,res := 1, ""
	for i := 0; i < len(frontStr); i++ {
		if i+1 < len(frontStr) {
			if frontStr[i] == frontStr[i+1] {
				sum++
			} else {
				res += strconv.Itoa(sum) + string(frontStr[i])
				sum = 1
			}
		} else {
			res += strconv.Itoa(sum) + string(frontStr[i])
		}
	}
	return res
}


func main() {
	fmt.Println(countAndSay(5))
}
