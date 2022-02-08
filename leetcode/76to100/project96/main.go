package main

import "fmt"

//输入：n = 3
//输出：5
func numTrees(n int) int {
	// 定义dp|n| 为以包含n的所有根节点所构成不同树的总和
	// dp|n| = f|1| + f|2| + ... + f|n|
	// 定义f|i| 为以i为根节点构成数的总和
	// f|i| = dp|i-1| * dp|n-i| 所有左子树与所有右子树的乘
	// 最终方程
	// dp|n| = dp|0| * dp|n-1| + dp|1| * dp|n-2| + ... + fp|n-1| * dp|0|

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
}
