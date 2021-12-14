package main

import "fmt"

//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
func minDistance(word1 string, word2 string) int {
	// dp|i|j| 表示A的前i个字母和B的前j个字母之间的编辑距离
	// len(word1)==0, dp|0|j| = len(word2), 即要操作len(word2)次才能使word1==word2,
	// 操作中可删除AB的字母或可添加AB字母
	n, m := len(word1), len(word2)
	if n*m == 0 {
		return n+m
	}
	dp := make([][]int, n)
	// 定义边界dp|0|0|=0
	for i := 0; i < n + 1; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = i
	}
	for i := 0; i < m + 1; i++ {
		dp[0][i] = i
	}
	for i:=1; i<n+1; i++ {
		for j:=1; j<m+1; j++ {

		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minDistance("intention", "execution"))
}
