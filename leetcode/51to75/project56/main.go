package main

import (
	"fmt"
	"sort"
)

//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	slow, fast, n := 0, 1, len(intervals)
	res := [][]int{}
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

func main()  {
	fmt.Println(merge([][]int{{1,4}, {2, 3}}))
}
