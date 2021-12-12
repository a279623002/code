package main

import "fmt"

//输入: nums = [1,3,5,6], target = 5
//输出: 2
func searchInsert(nums []int, target int) int {
	left, right, res := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := (left + right)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid-1
			res = mid
		} else {
			left = mid+1
		}
	}
	return res
}

func main() {
	fmt.Println(searchInsert([]int{1}, 0))
}
