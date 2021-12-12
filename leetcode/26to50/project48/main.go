package main

import "fmt"

//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[[7,4,1],[8,5,2],[9,6,3]]
func rotate(matrix [][]int)  {
	n := len(matrix)
	// 789
	// 456
	// 123
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	// 对角线翻转
	for i := 0; i < n-1; i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {
	demo := [][]int{{1, 2, 3},{4,5,6},{7,8,9}}
	rotate(demo)
	fmt.Println(demo)
}
