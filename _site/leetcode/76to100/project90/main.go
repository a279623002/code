package main

import (
	"fmt"
	"sort"
)

//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	carry := []int{}

	var fn func(bool,int)
	fn = func(ok bool, i int) {
		if i == len(nums) {
			res = append(res, append([]int(nil), carry...))
		} else {
			fn(false, i+1)
			// ok 上一个没有使用，即树层
			// ok 要是上一个使用，即树枝
			if !ok && i > 0 && nums[i-1] == nums[i] {
				return
			}
			carry = append(carry, nums[i])
			fn(true, i+1)
			carry = carry[:len(carry)-1]
		}
	}
	fn(false, 0)
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
