package main

import "fmt"

//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		if vv, ok := m[v]; ok {
			res = []int{k, vv}
		}
		m[target - v] = k
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}