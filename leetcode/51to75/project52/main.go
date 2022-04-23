package main

import "fmt"

//输入：n = 4
//输出：2
//解释：如上图所示，4 皇后问题存在两个不同的解法。
func totalNQueens(n int) int {
	res := 0
	columns, l, r := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	var fn func(int)
	fn = func(row int) {
		if row == n {
			res++
		} else {
			for i := 0; i < n; i++ {
				if columns[i] {
					continue
				}
				if l[row-i] {
					continue
				}
				if r[row+i] {
					continue
				}
				columns[i], l[row-i], r[row+i] = true, true, true
				fn(row+1)
				columns[i], l[row-i], r[row+i] = false, false, false
			}
		}
	}
	fn(0)
	return res
}

func main() {
	fmt.Println(totalNQueens(4))
}
