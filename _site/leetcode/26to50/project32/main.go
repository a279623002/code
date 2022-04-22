package main

import "fmt"

//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
func longestValidParentheses(s string) int {
	// dp|i| = dp|i-2| + 2 						s[i-1] == '(' && s[i] == ')'
	// 		or dp|i-1| + dp|i-dp|i-1|-2| + 2 	s[i-1] == ')' && s[i] == ')' && s[i - dp|i-1|-2 + 1] == '('
	dp := map[int]int{}
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// ....()
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i - dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				// ....))
				if i - dp[i-1] >=2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	fmt.Println(longestValidParentheses("()(()()")) // ()() 4
}

