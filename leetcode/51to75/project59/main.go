package main

import "fmt"

//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
func generateMatrix(n int) [][]int {
	row, column := n, n
	res := make([][]int, row)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, column)
	}

	// 上： (top, left), (top, right)
	// 右： (top+1, right), (bottom-1, right)
	// 下： (bottom, right), (bottom, left)
	// 左： (bottom-1, left), (top+1, left)
	count, top, right, bottom, left := 1, 0, n-1, n-1, 0
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			res[top][col] = count
			count++
		}
		for row := top+1; row <= bottom-1; row++ {
			res[row][right] = count
			count++
		}
		// 需要加层判断，不然会重复计算
		if left < right && top < bottom {
			for col := right; col >= left; col-- {
				res[bottom][col] = count
				count++
			}
			for row := bottom-1; row >= top+1; row-- {
				res[row][left] = count
				count++
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return res
}

func main() {
	fmt.Println(generateMatrix(3))
}
