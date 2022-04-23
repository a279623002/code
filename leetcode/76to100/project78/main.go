package main

import "fmt"

//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets(nums []int) [][]int {
	res := [][]int{}
	carry := []int{}
	var fn func(int)
	fn = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int(nil), carry...))
			return
		}
		carry = append(carry, nums[cur])
		fn(cur+1)
		carry = carry[:len(carry)-1]
		fn(cur+1)
	}
	fn(0)
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
