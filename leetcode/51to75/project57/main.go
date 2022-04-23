package main

import (
	"fmt"
	"sort"
)

//输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
//输出：[[1,5],[6,9]]
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	slow, fast, n := 0, 1, len(intervals)
	for fast < n {
		if intervals[slow][0] <= intervals[fast][0] && intervals[fast][0] <= intervals[slow][1] {
			if intervals[slow][1] < intervals[fast][1] {
				intervals[slow][1] = intervals[fast][1]
			}
		} else {
			res = append(res, intervals[slow])
			slow = fast
		}
		fast++
	}
	res = append(res, intervals[slow])
	return res
}

func main() {
	fmt.Println(insert([][]int{{1,3}, {6, 9}}, []int{2,5}))
}
