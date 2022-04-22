package main

import "fmt"

//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
func jump(nums []int) int {
	// 能跳的范围内取最大值，如果达到最后，就结束
	count, n := 0, len(nums)
	end, start := 0, nums[0]
	for i := 0; i < n-1; i++ {
		// i是0到num[i],加上num[i]就是下一个下一次能跳最长的值
		start = max(start, i+nums[i])
		if i == end {
			// 到达位置时更新下一次最长的值，同时增加次数，最后一个位置不跳所以只遍历n-1
			end = start
			count++
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(jump([]int{2,3,0,1,4}))
}