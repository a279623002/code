package main

import "fmt"

//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length % 2 == 1 {
		// 奇数
		middle := length / 2
		return float64(getKthElement(nums1, nums2, middle + 1))
	} else {
		// 偶数
		middle1, middle2 := length / 2 - 1, length / 2
		return float64(getKthElement(nums1, nums2, middle1 + 1) + getKthElement(nums1, nums2, middle2 + 1))/2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		// 当下标等于该长度时，既已经全部排除，返回另一个的中间值
		if index1 == len(nums1) {
			return nums2[index2 + k - 1]
		}
		if index2 == len(nums2) {
			return nums1[index1 + k - 1]
		}
		// 两个数组以排除到最后一个，返回最小那个
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		// 取两个数组对应的中间值，如果超出长度，取最后一个，可以一次性排除
		half := k / 2
		newIndex1 := min(index1 + half, len(nums1)) - 1
		newIndex2 := min(index2 + half, len(nums2)) - 1
		if nums1[newIndex1] <= nums2[newIndex2] {
			// 防止越界k -= k/2 - 1 换成下面方式依靠排除的个数减少k值
			k -= (newIndex1 - index1 + 1)
			// 移动下标--排除前面
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
