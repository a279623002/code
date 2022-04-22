package main

//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for ri, row := range board {
		for ci, col := range row {
			// col转数字
			if col == '.' {
				continue
			}
			num := col - '1'
			rows[ri][num]++
			columns[ci][num]++
			subboxes[ci/3][ri/3][num]++
			if rows[ri][num] > 1 || columns[ci][num] > 1 || subboxes[ci/3][ri/3][num] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {

}
