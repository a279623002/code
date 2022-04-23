package main

import "fmt"

//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]
func sortColors(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	start, left, right := nums[0], 0, len(nums)-1
	i := 1
	for left < right {
		if nums[start] < nums[i] {
			// nums[i]放到右边，right-1，继续遍历当前nums[i]
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			// nums[i]放到左边，left+1，进入下一个遍历i+1
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		}
	}
	// 排序后递归，直到nums需要排序的个数小于或等于1
	// 		左边--				start				--右边
	// 左边--start--右边		左边--start--右边		左边--start--右边
	sortColors(nums[:left])
	sortColors(nums[right+1:])
}

func main() {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
	fmt.Println(nums)
}
