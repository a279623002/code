package main

import "fmt"

//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	var fn func([]int)
	fn = func(tmp []int) {
		if len(tmp) == n {
			res = append(res, tmp)
		} else {
			for i:=0; i<n; i++ {
				// 跳过在tmp里的元素
				skip := false
				for j:=0; j<len(tmp); j++ {
					if nums[i] == tmp[j] {
						skip = true
						break
					}
				}
				if !skip {
					fn(append(tmp, nums[i]))
				}
			}

		}
	}
	fn([]int{})
	return res
}

func permute1(nums []int) [][]int {
	res := [][]int{}
	var fn func([]int, []int)
	fn = func(n []int, tmp []int) {
		if len(n) == 0 {
			res = append(res, tmp)
		} else {
			length := len(n)
			for i:=0; i<length; i++ {
				carry := n[i]
				// 剔除元素
				newNum := make([]int, len(n))
				// 复制切片，防止修改到原来的切片
				copy(newNum, n)
				tmpNums := []int{}
				if length > 1 {
					if i == 0 {
						tmpNums = newNum[1:]
					} else if i == length-1 {
						tmpNums = newNum[:i]
					} else {
						tmpNums = append(newNum[:i], newNum[i+1:]...)
					}
				}
				fn(tmpNums, append(tmp, carry))
			}
		}
	}
	fn(nums, []int{})
	return res
}

func permute2(nums []int) [][]int {
	res := [][]int{}
	tmp := []int{}
	vis := make([]bool, len(nums))
	var fn func(int)
	fn = func(n int) {
		if n == len(nums) {
			res = append(res, append([]int(nil), tmp...))
		} else {
			for k, v := range nums {
				// 已记录则跳过
				if vis[k] {
					continue
				}
				vis[k] = true
				tmp = append(tmp, v)
				fn(n+1)
				tmp = tmp[:len(tmp)-1]
				vis[k] = false
			}
		}
	}
	fn(0)
	return res
}

func main() {
	fmt.Println(permute2([]int{1,1,2}))
}
