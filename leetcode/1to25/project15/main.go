package main

import (
	"fmt"
	"sort"
)

//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	res := [][]int{}
	// 枚举a
	for first := 0; first < len(nums); first++ {
		// 第二次枚举a的时候要跟上一次的数不同
		if first > 0 && nums[first] == nums[first-1] {
			continue 
		}
		// 右指针
		third := len(nums) - 1
		// 枚举b 左指针
		for second := first + 1; second < len(nums); second++ {
			// 第二次枚举b的时候要跟上一次的数不同
			if second > first + 1 && nums[second] == nums[second-1] {
				continue
			}
			// 保证右指针在左指针右边，3数和大于0时不断缩进
			for third > second && nums[first] + nums[second] + nums[third] > 0 {
				third--
			}
			// 左右指针不能相等
			if second == third {
				break
			}
			if nums[first] + nums[second] + nums[third] == 0 {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}