package main

import (
	"fmt"
)

//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true
func exist(board [][]byte, word string) bool {
	r, c := len(board), len(board[0])
	// 记录已访问过的
	visit := make([][]bool, r)
	for i := range visit {
		visit[i] = make([]bool, c)
	}

	// row, col, word的index
	var fn func(int, int, int) bool
	fn = func(row int, col int, index int) bool {
		// 剪枝
		if board[row][col] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		visit[row][col] = true
		// 用defer回溯，后面有好多个return
		defer func() { visit[row][col] = false }()
		// 上下左右搜索
		if row-1 >= 0 && !visit[row-1][col] {
			if fn(row-1, col, index+1) {
				return true
			}
		}
		if row+1 < r && !visit[row+1][col] {
			if fn(row+1, col, index+1) {
				return true
			}
		}
		if col-1 >= 0 && !visit[row][col-1] {
			if fn(row, col-1, index+1) {
				return true
			}
		}
		if col+1 < c && !visit[row][col+1] {
			if fn(row, col+1, index+1) {
				return true
			}
		}
		return false
	}
	for rr, v := range board {
		for cc := range v {
			if fn(rr, cc, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	//fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	fmt.Println(exist([][]byte{{'a','b'}}, "ab"))
}
