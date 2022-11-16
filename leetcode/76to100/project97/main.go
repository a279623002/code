package main

import "fmt"

// 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
//
// 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
//
// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// |n - m| <= 1
// 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
// 注意：a + b 意味着字符串 a 和 b 连接
// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// 输出：true
// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// 输出：false
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
