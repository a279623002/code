package main

import "fmt"

//输入：nums = [2,5,6,0,0,1,2], target = 0
//输出：true
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return true
		}
		// 防止找不到升序的一边
		// 1,0,1 	1, 1, 1
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else {
			if nums[left] <= nums[mid] {
				// 左边升序
				if nums[left] <= target && target < nums[mid] {
					right = mid-1
				} else {
					left = mid+1
				}
			} else {
				// 右边升序
				if nums[mid] < target && target <= nums[right] {
					left = mid+1
				} else {
					right = mid-1
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(search([]int{3, 1}, 1))
	//fmt.Println(search([]int{2,5,6,0,0,1,2}, 3))
}
