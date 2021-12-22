package main

import (
	"fmt"
	"sort"
)

//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//解释：需要合并 [1,2,3] 和 [2,5,6] 。
//合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}
func merge1(nums1 []int, m int, nums2 []int, n int) {
	arr := []int{}
	i, j := 0, 0
	for i < m || j < n {
		if i >= m && j < n {
			arr = append(arr, nums2[j])
			j++
		} else if j >= n && i < m {
			arr = append(arr, nums1[i])
			i++
		} else if j < n && i < m {
			if nums1[i] < nums2[j] {
				arr = append(arr, nums1[i])
				i++
			} else {
				arr = append(arr, nums2[j])
				j++
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		nums1[i] = arr[i]
	}
}

func main() {
	nums1 := []int{7, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 1, nums2, 3)
	fmt.Println(nums1)
}
