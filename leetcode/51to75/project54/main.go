package main

import "fmt"

//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
func spiralOrder(matrix [][]int) []int {
	row, column := len(matrix), len(matrix[0])
	res := make([]int, row*column)
	index := 0
	// 上 (top, left)--(top, right)
	// 右 (top+1, right)--(bottom, right)
	// 下 (bottom, right-1)--(bottom, left-1)
	// 左 (bottom, left)--(top-1, left)
	top, bottom, left, right := 0, row-1, 0, column-1
	for left <= right && top <= bottom {
		// 加上层
		for col := left; col <= right; col++ {
			res[index] = matrix[top][col]
			index++
		}
		// 加右层
		for row := top + 1; row <= bottom; row++ {
			res[index] = matrix[row][right]
			index++
		}
		// 需要加层判断，不然会重复计算
		if left < right && top < bottom {
			// 加下层
			for col := right - 1; col > left; col-- {
				res[index] = matrix[bottom][col]
				index++
			}
			// 加左层
			for row := bottom; row > top; row-- {
				res[index] = matrix[row][left]
				index++
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
