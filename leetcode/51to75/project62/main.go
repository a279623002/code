package main

import "fmt"

//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下

// 动态规划
func uniquePaths(m int, n int) int {
	//dp|i|j| = dp|i-1|j| + dp|i|j-1| 左上角走到(i,j) 的路径数量
	//初始条件为 dp|0|0|=1，即从左上角走到左上角有一种方法
	// dp|0|j| dp|i|0| 都设为边界条件
	// 1 1 1 1
	// 1 2 3 4
	// 1 3 6 10
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
// 回溯超时
func uniquePaths1(m int, n int) int {
	res := 0
	if m == 1 && n == 1 {
		return 1
	}
	if m > 0 {
		res += uniquePaths(m-1, n)
	}
	if n > 0 {
		res += uniquePaths(m, n-1)
	}
	return res
}

func main() {
	fmt.Println(uniquePaths(4, 3))
}
