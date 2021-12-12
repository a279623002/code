package main

import "fmt"

//输入：s = "PAYPALISHIRING", numRows = 4
//输出："PINALSIGYAHRPI"
//解释：
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var res string
	// 主列上对应的索引
	skip := (numRows - 1) * 2
	// 按行读取
	for i := 0; i < numRows; i++ {
		for j := 0; j + i < len(s); j += skip {
			res += string(s[i + j])
			// 首行和尾行没有中间列
			if i != 0 && i != numRows - 1 && j + skip - i < len(s) {
				// 中间列对应的索引
				res += string(s[j + skip - i])
			}
		}
	}
	return res
}

func main() {
	fmt.Println(convert("PAYPALISHIRING",4))
}