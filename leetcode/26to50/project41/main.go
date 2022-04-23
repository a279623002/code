package main

import "fmt"

//输入：nums = [3,4,-1,1]
//输出：2
func firstMissingPositive(nums []int) int {
	// 置换数组为如下形式，nums[x-1]=x 即 nums[3-1]=3 (1,2,3,4)
	// 遍历0到N（原数组长度）
	// 	如果出现nums[x-1]!=x 即nums[2-1] != 2 (1,-1,2,3) 就返回x
	//	如果每个位置都正确, 说明0到N都有正确的x，返回N+1 (1,2,3,4)=》返回5
	n := len(nums)
	for i := 0; i < n; i++ {
		// 每个位置要遍历到无法值换才进入到下一次遍历
		// 4123=>3124=>2134=>1234
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i+1
		}
	}
	return n+1
}

func main() {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}
