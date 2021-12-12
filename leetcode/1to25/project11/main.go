package main

import "fmt"

//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
func maxArea(height []int) int {
	left, right, res := 0, len(height) - 1, 0
	for left < right {
		l, r := height[left], height[right]
		res = max(res, min(l, r) * (right - left))
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
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}