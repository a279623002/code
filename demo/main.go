package main

import (
	"fmt"
)

func fn(nums []int) int {
	left, right, res := 0, len(nums)-1, 0
	for left < right {
		l, r := nums[left], nums[right]
		res = max(res, min(l, r) * (right-left+1))
		if l > r {
			right--
		} else {
			left++
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(fn([]int{2, 3, 3, 1}))
}
