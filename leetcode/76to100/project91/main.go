package main

import "fmt"

//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//输入：s = "12"
//输出：2
//解释：它可以解码为 "AB"（1 2）或者 "L"（12）。
func numDecodings(s string) int {
	// dp|i| = dp|i-1| + dp|i-2| (nums[i-2] != 0 && nums[i-1] != 0 && nums[i-2:i]<27)
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i:=1; i<=n; i++ {
		if s[i-1]!= '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') < 27 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func main() {
	// 1 2 3 4 5
	// 12  3 4 5
	// 1 23  4 5
	//
	fmt.Println(numDecodings("3123"))
}
