package main

import (
	"fmt"
	"sort"
)

//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
//[1,2,1],
//[2,1,1]]
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	tmp := []int{}
	vis := make([]bool, len(nums))
	var fn func(int)
	fn = func(n int) {
		if n == len(nums) {
			res = append(res, append([]int(nil), tmp...))
		} else {
			for k, v := range nums {
				// 已记录则跳过，跳过已填入的重复值
				if vis[k] || k > 0 && !vis[k-1] && v == nums[k-1] {
					continue
				}
				vis[k] = true
				tmp = append(tmp, v)
				fn(n+1)
				tmp = tmp[:len(tmp)-1]
				vis[k] = false
			}
		}
	}
	fn(0)
	return res
}
func main() {
	fmt.Println(permuteUnique([]int{1,1,2}))
}