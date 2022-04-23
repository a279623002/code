package main

import (
	"fmt"
	"sort"
)

//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for first := 0; first < len(nums); first++ {
		// 如果大于0即有两个或以上的first，跳过重复
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		for ; second < len(nums); second++ {
			// 如果大于first + 1即有两个或以上的second，跳过重复
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			third := second + 1
			for ; third < len(nums); third++ {
				// 如果大于second + 1即有两个或以上的third，跳过重复
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}
				fourth := len(nums)-1
				for third < fourth && nums[first] + nums[second] + nums[third] + nums[fourth] != target {
					fourth--
				}
				if third == fourth {
					continue
				}

				res = append(res, []int{nums[first], nums[second], nums[third], nums[fourth]})
			}
		}
	}
	return res
}

func main() {
	//-2 -1 0 0 1 2
	fmt.Println(fourSum([]int{1,0,-1,0,-2,2}, 0))
}
