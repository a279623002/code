package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	// 定义dp|i||j| 为 s1的前i个字符与s2的前j个字符能否交错组成s3
	// dp|i|j| = dp|i-1|j| (s1[i-1] == s3[i+j-1] || dp|i|j-1| (s2[j-1] == s3[i+j-1])
	n, m := len(s1), len(s2)
	if len(s3) != n+m {
		return false
	}
	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	// 处理边界值
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for i := 1; i <= m; i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1] || dp[i][j-1] && s2[j-1] == s3[i+j-1]
		}
	}
	return dp[n][m]
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}
