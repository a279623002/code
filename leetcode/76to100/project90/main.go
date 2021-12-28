package main

import (
	"fmt"
	"sort"
)

//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
func subsetsWithDup(nums []int) [][]int {
	//		1		2		(2)
	//	2	(2)		2
	// 2
	sort.Ints(nums)
	res := [][]int{}
	tmp := []int{}
	var fn func(bool, int)
	fn = func(ok bool, i int) {
		if i == len(nums) {
			res = append(res, append([]int(nil), tmp...))
		} else {
			fn(false, i+1)
			//要去重的是同一树层上的“使用过”，同一树枝上的都是一个组合里的元素，不用去重。
			if !ok && i > 0 && nums[i-1] == nums[i] {
				return
			}

			tmp = append(tmp, nums[i])
			fn(true, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	fn(false, 0)
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
