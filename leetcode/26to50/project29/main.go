package main

import (
	"fmt"
	"math"
)

//输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend // 只要不是最小值，返回相反数
		}
		return math.MaxInt32 // 最小值就返回最大值
	}
	//统一用正数，并且记录正负
	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	res := div(dividend, divisor)
	if sign != 1 {
		res = -res
	}
	return res
}

// 60/8 = (60 - 32)/8 + 4 = (60 - 32 - 16)/8 + 4 + 2 =  (60 - 32 - 16 - 8)/8 + 4 + 2 + 1 = 7
func div(dividend, divisor int) int {
	if dividend < divisor {
		return 0
	}
	count := 1
	tmp := divisor
	for tmp + tmp < dividend {
		count += count
		tmp += tmp
	}
	return count + div(dividend - tmp, divisor)
}

func main() {
	fmt.Println(divide(60, 8))
}
