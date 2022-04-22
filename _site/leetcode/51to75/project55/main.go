package main

import "fmt"

//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
func canJump(nums []int) bool {
	m, n := nums[0], len(nums)
	for i := 0; i < n; i++ {
		// 当前能跳最大的位置到达不了i，直接返回false
		if m < i {
			return false
		}
		// 更新最大位置
		m = max(m, i+nums[i])
		// 到达终点直接返回
		if m >= n-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(canJump([]int{0}))
}
