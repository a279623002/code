package main

import "fmt"

//输入：nums = [1,2,3]
//输出：[1,3,2]
// 12345678
// 12345687
// 12345768
// 12345786
// 12346578
// 12346758
// ...
// 87654321
// 升序最小
// 1.下一个比当前大，即后面大的数与前面小的数交换，如12345678 =》 12345687
// 2. 增加的幅度尽可能小
//		2.1 从后往前面扫，交换尽可能小的大数和前面的小数， 如12345687 ，下一个排列应该是7跟6交换而不是6跟5
//		2.2 交换后需要把后面的重置为升序，得到下一个紧靠的大数
// 1. 从后向前查找第一个相邻升序的元素对(i, j)，满足A[i]<A[j]，此时[j,end)一定是降序
// 2. 在[j,end)从后往前找到第一个满足A[i]<A[k], A[i]就是2.1中的小数, A[K]则是大数
// 3. 交换A[i],A[k],此时[j,end)一定是降序
// 4. 重置[j,end)为升序
func nextPermutation(nums []int)  {
	numsLen := len(nums)
	if numsLen <= 1 {
		return
	}

	ai, aj, ak := numsLen-2, numsLen-1, numsLen-1
	for ai >= 0 && nums[ai] >= nums[aj] {
		ai--
		aj--
	}

	if ai >= 0 {
		for nums[ai] >= nums[ak] {
			ak--
		}
		nums[ai], nums[ak] = nums[ak], nums[ai]
	}

	for i, j := aj,len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println(nums)
}

func main() {
	nextPermutation([]int{3,2,1})
}
