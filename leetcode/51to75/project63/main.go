package main

import "fmt"

//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右

//dp[i][j]=dp[i-1][j]+dp[i][j-1] => dp[j] = dp[j] + dp[j-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	// 使用滚动数组
	// 如遍历到A时 a[0] = 1, a[j-1] = 1, a[j] = 1
	// 下一个遍历B a[0] = 1, a[j-1] = 2, a[j] = 3
	// B[j] = B[j] + B[j-1] ==> A[j] + B[j-1]
	arr := make([]int, m)
	// 初始化为1，后面遍历用-1获取
	arr[0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				arr[j] = 0
				continue
			}
			if j > 0 {
				arr[j] += arr[j-1]
			}
		}
	}
	return arr[m-1]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	// 1 1 1 1
	// 1 0 1 2
	// 1 1 2 4
	i, j := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, i)
	for k, _ := range dp {
		dp[k] = make([]int, j)
	}
	if obstacleGrid[0][0] != 1 && obstacleGrid[i-1][j-1] != 1 {
		dp[0][0] = 1
		for n := 0; n < i; n++ {
			for m := 0; m < j; m++ {
				if n == 0 || m == 0 {
					if m > 0 {
						if obstacleGrid[n][m-1] == 0 && obstacleGrid[n][m] == 0 {
							dp[n][m] = dp[n][m-1]
						}
					}
					if n > 0 {
						if obstacleGrid[n-1][m] == 0 && obstacleGrid[n][m] == 0 {
							dp[n][m] = dp[n-1][m]
						}
					}
				} else {
					if obstacleGrid[n-1][m] == 0 && obstacleGrid[n][m] == 0  {
						dp[n][m] += dp[n-1][m]
					}
					if obstacleGrid[n][m-1] == 0 && obstacleGrid[n][m] == 0  {
						dp[n][m] += dp[n][m-1]
					}
				}
			}
		}
	}
	// 0 0 0      1 0 0
	// 0 1 0      0 0 0
	// 0 0 0      0 0 0
	fmt.Println(dp)
	return dp[i-1][j-1]
}

// 回溯超时
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	count := 0
	var fn func(int, int)
	fn = func(n int, m int) {
		if n == 0 && m == 0 {
			count++
		} else {
			if n < 0 || m < 0 {
				return
			}
			if obstacleGrid[n][m] == 1 {
				return
			}

			fn(n-1, m)
			fn(n, m-1)
		}
	}
	n, m := len(obstacleGrid)-1, len(obstacleGrid[0])-1
	if obstacleGrid[0][0] != 1 && obstacleGrid[n][m] != 1 {
		fn(n, m)
	}
	return count
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}
