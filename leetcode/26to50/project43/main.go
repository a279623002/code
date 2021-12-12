package main

import (
	"fmt"
	"strconv"
)

//输入: num1 = "123", num2 = "456"
//输出: "56088"
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := ""
	for i := len(num1) - 1; i >= 0; i-- {
		cur, carry := "", 0
		for k := len(num1) - 1; k > i; k-- {
			// 补0
			cur += "0"
		}
		x := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			sum := x * y + carry
			// 取出当前位数，如12拿2，1存到carry后面再加上去
			cur = strconv.Itoa(sum%10) + cur
			carry = sum/10
		}

		for ; carry != 0; carry/=10 {
			cur = strconv.Itoa(carry%10) + cur
		}
		res = addString(res, cur)
	}
	return res
}

func addString(num1 string, num2 string) string {
	res, carry := "", 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		res = strconv.Itoa(sum) + res
	}
	//if carry > 0 {
	//	res = strconv.Itoa(carry) + res
	//}
	return res
}

func main() {
	fmt.Println(multiply("123", "4"))
}
