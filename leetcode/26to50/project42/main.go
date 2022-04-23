package main

import "fmt"

//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
func trap(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	// 每个下标先存下左右的最大高度，累加最小的高度减去当前值
	n := len(height)
	lefts, rights := make([]int, n), make([]int, n)
	lefts[0] = height[0]
	for i := 1; i < n; i++ {
		lefts[i] = max(height[i], lefts[i-1])
	}
	rights[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rights[i] = max(height[i], rights[i+1])
	}

	res := 0
	for i := 1; i < n-1; i++ {
		//fmt.Println(rights[i], lefts[i], height[i])
		res += min(rights[i], lefts[i]) - height[i]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}
