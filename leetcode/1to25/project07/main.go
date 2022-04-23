package main

import (
	"fmt"
	"math"
)

//输入：x = 123
//输出：321
func reverse(x int) int {
	res := 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		res = res * 10 + x % 10
		x /= 10
	}
	return res
}

func main() {
	//fmt.Println(math.MaxInt32)
	//fmt.Println((1<<31) - 1)
	//fmt.Println(math.MinInt32)
	//fmt.Println(-(1<<31))
	fmt.Println(reverse(-1212))
}
