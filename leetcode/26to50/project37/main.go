package main

//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
func solveSudoku(board [][]byte) {
	var rows, columns [9][9]bool
	var block [3][3][9]bool
	var space [][2]int
	// 取出要填的空格下标，并且更新行列的规则
	for ri, row := range board {
		for ci, col := range row {
			if col == '.' {
				space = append(space, [2]int{ri, ci})
			} else {
				num := col - '1'
				rows[ri][num] = true      //该行已有这个数字
				columns[ci][num] = true   //该列已有这个数字
				block[ri/3][ci/3][num] = true //该区域已有这个数字
			}
		}
	}

	// 递归方式填数字
	var fn func(int) bool
	fn = func(pos int) bool {
		if pos == len(space) {
			return true
		}
		ri, ci := space[pos][0], space[pos][1]

		for num := byte(0); num < 9; num++ {
			if (!rows[ri][num] && !columns[ci][num] && !block[ri/3][ci/3][num]) {
				// 填入数字时更新规则
				rows[ri][num] = true
				columns[ci][num] = true
				block[ri/3][ci/3][num] = true
				board[ri][ci] = num + '1'
				// 递归填入
				if (fn(pos+1)) {
					return true
				}

				// 在回溯到当前递归层时，即数字不符合条件 重置规则，进入下一个遍历
				rows[ri][num] = false
				columns[ci][num] = false
				block[ri/3][ci/3][num] = false
			}
		}
		return false
	}
	fn(0)
}

func main() {

}
