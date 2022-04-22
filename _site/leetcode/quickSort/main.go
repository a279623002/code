package main

import "fmt"

func quick(nums []int) {
	if len(nums) <= 1 {
		return
	}
	start, left, right := nums[0], 0, len(nums)-1
	i := 1
	for left < right {
		if nums[i] > start {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		} else {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		}
	}
	quick(nums[:left])
	quick(nums[right+1:])

}

func main() {
	nums := []int{7, 4, 3, 2, 5, 8, 1, 9, 0, 6}
	quick(nums)
	fmt.Println(nums)
}
