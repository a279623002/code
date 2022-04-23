package main

import (
	"fmt"
	"sort"
)

//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	comb := []int{}
	//freq := [][2]int{} // 0存值，1存该值的次数
	var freq [][2]int
	for _, num := range candidates {
		// 已经排序，不是重复值就新加freq
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}
	fmt.Println(freq)

	var fn func(int, int)
	fn = func(target, idx int) {
		if target == 0 {
			res = append(res, append([]int(nil), comb...))
			return
		}
		if idx == len(freq) || target < freq[idx][0] {
			return
		}
		fn(target, idx+1) // 取单个直接命中结果

		most := min(target/freq[idx][0], freq[idx][1])
		for i:=1; i <= most; i++ {
			comb = append(comb, freq[idx][0])
			fn(target-freq[idx][0]*i, idx+1)
		}
		comb = comb[:len(comb)-most]

	}
	fn(target, 0)
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(combinationSum2([]int{2,5,2,1,2}, 5))
}
