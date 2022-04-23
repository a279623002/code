package main

import (
	"fmt"
	"sort"
)

//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
func searchMatrix(matrix [][]int, target int) bool {
	// 使用内置方法
	n, m := len(matrix), len(matrix[0])
	row := sort.Search(n, func(i int) bool {
		return matrix[i][0] > target
	})
	if row < 1 {
		return false
	}
	//col := sort.Search(m, func(i int) bool {
	//	return matrix[row-1][i] > target
	//})
	//return col <= m && matrix[row-1][col-1] == target

	col := sort.SearchInts(matrix[row-1], target)

	return col <= m && matrix[row-1][col] == target

}

func searchMatrix1(matrix [][]int, target int) bool {
	// 先对row进行查找,然后在对col查找
	n, m := len(matrix), len(matrix[0])
	rowLeft, rowRight := 0, n-1
	for rowLeft <= rowRight {
		rowMid := (rowLeft+rowRight)/2
		if matrix[rowMid][0] > target {
			rowRight = rowMid-1
		} else if matrix[rowMid][0] < target {
			rowLeft = rowMid+1
		} else {
			return true
		}
	}
	if rowRight < 0 {
		// matrix[0][0] < target
		return false
	}
	colLeft, colRight := 0, m-1
	for colLeft <= colRight {
		colMid := (colLeft+colRight)/2
		if matrix[rowRight][colMid] > target {
			colRight = colMid-1
		} else if matrix[rowRight][colMid] < target {
			colLeft = colMid+1
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1,3,5,7}, {10,11,16,20}, {23,30,34,60}},60))
}
