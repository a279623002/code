package main

import "fmt"

//输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
//输出：[[1,0,1],[0,0,0],[1,0,1]]
func setZeroes(matrix [][]int)  {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for r, v := range matrix {
		for c, vv := range v {
			if vv == 0 {
				row[r] = true
				col[c] = true
			}
		}
	}
	for r, v := range matrix {
		for c, _ := range v {
			if row[r] || col[c] {
				matrix[r][c] = 0
			}
		}
	}
}

func main() {
	nums := [][]int{{1,1,1}, {1,0,1}, {1,1,1}}
	setZeroes(nums)
	fmt.Println(nums)
}
