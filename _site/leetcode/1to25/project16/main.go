package main

import (
	"fmt"
	"math"
	"sort"
)

//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := math.MaxInt32
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second, third := first+1, len(nums)-1
		for second < third {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			sum := nums[first] + nums[second] + nums[third]
			if sum == target {
				return target
			}

			res = min(res, sum, target)
			if sum > target {
				t := third - 1
				for second < t && nums[t] == nums[third] {
					t--
				}
				third = t
			} else {
				s := second + 1
				for s < third && nums[second] == nums[s] {
					s++
				}
				second = s
			}
		}
	}
	return res
}

func min(a, b, target int) int {
	if abs(target - a) < abs(target - b) {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func main() {
	fmt.Println(threeSumClosest([]int{1, 2, 5, 10, 11}, 12))
}
