package main

import "fmt"

//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//皇后的走法是：可以横直斜走，格数不限。
//因此要求皇后彼此之间不能相互攻击，等价于要求任何两个皇后都不能在同一行、同一列以及同一条斜线上
func solveNQueens(n int) [][]string {
	res := [][]string{}
	// 每一行queen存在的位置
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	// 同列集合, 左上到右下, 右上到左下斜线集合
	columns, l, r := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	var fn func([]int, int)
	fn = func(queens []int, row int) {
		if row == n {
			board := getBoard(queens, n)
			res = append(res, board)
		} else {
			for i := 0; i < n; i++ {
				if columns[i] {
					continue
				}
				// 左上到右下标记的斜线，行下标-列下标的值是相等的
				if l[row-i] {
					continue
				}
				// 右上到左下标记的斜线，行下标+列下标的值是相等的
				if r[row+i] {
					continue
				}
				columns[i], l[row-i], r[row+i] = true, true, true
				queens[row] = i
				fn(queens, row+1)
				queens[row] = -1
				columns[i], l[row-i], r[row+i] = false, false, false
			}
		}
	}
	fn(queens, 0)
	return res
}

// 生成棋盘
func getBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

func main() {
	fmt.Println(solveNQueens(1))
}
