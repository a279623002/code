package main

import "fmt"

//输入: candidates = [2,3,6,7], target = 7
//输出: [[7],[2,2,3]]
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	comb := []int{}
	var fn func(int, int)
	fn = func(target int, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int(nil), comb...))
			return
		}
		// 跳过
		fn(target, idx+1)
		// 重复使用
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			fn(target-candidates[idx], idx)
			//fmt.Println(comb)
			comb = comb[:len(comb)-1]
		}
	}
	fn(target, 0)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}
