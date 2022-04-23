package main

import (
	"fmt"
	"strconv"
)

//"123"
//"132"
//"213"
//"231"
//"312"
//"321"
//输入：n = 3, k = 3
//输出："213"
func getPermutation(n int, k int) string {
	// 阶乘
	factorial := make([]int, n) // 1 1 2 6 24 120 720 5040 40320
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1]*i
	}
	fmt.Println(factorial)
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		nums[i] = i
	}
	// k为名次，减去1才是值，然后逆推
	k -= 1
	res := ""
	for i := n; i >= 1; i-- {
		mod := k%factorial[i-1]
		index := k/factorial[i-1]
		k = mod
		res += strconv.Itoa(nums[index+1])
		// 弹出去index+1，后续index根据k值在剩下的值寻找
		if index+1 == len(nums)-1 {
			nums = nums[:index+1]
		} else {
			nums = append(nums[:index+1], nums[index+2:]...)
		}
	}
	return res
}

func main() {
	fmt.Println(getPermutation(9, 40320*9))
}