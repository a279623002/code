package main

import (
	"fmt"
)

//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
func searchRange(nums []int, target int) []int {
	l := binary(nums, target, true)
	r := binary(nums, target, false) - 1

	if l <= r && r < len(nums) && nums[l] == target &&  nums[r] == target {
		return []int{l, r}
	}
	return []int{-1, -1}
}

// 该二分法不是直接命中，是查找到最后left<=right才结束，中间命中条件时会记录mid，更新到最后命中的结果
// lower为ture则最后命中结果大于等于target false则是大于target
func binary(nums []int, target int, lower bool) int {
	left, right, res := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := (left + right)/2
		if nums[mid] > target || (nums[mid] >= target && lower) {
			right = mid-1
			res = mid
		} else {
			left = mid+1
		}
	}
	return res
}


func main() {
	fmt.Println(searchRange([]int{}, 8))
}