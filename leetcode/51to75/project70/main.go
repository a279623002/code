package main

import "fmt"

//输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
func climbStairs(n int) int {
	// dp|i| = dp|i-1| + dp|i-2|
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] += dp[i-1]
		if i > 1 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
func climbStairs1(n int) int {
	// 0 =》 1
	// 1 =》 1
	// 2 =》 2
	// 3 =》 3
	// 4 =》 5
	// 5 =》 8
	// dp|i| = dp|i-1| + dp|i-2|
	o, t, res := 0, 0, 1
	for i := 0; i < n; i++ {
		// 滚动数组
		o = t
		t = res
		res = o + t
	}
	return res
}

func main() {
	fmt.Println(climbStairs(44))
}
