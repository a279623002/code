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

//注意：为什么我们总是在单词 A 和 B 的末尾插入或者修改字符，能不能在其它的地方进行操作呢？
// 答案是可以的，但是我们知道，操作的顺序是不影响最终的结果的。
// 例如对于单词 cat，我们希望在 c 和 a 之间添加字符 d 并且将字符 t 修改为字符 b，
// 那么这两个操作无论为什么顺序，都会得到最终的结果 cdab
// 就不要想怎么增删改，当成一次操作次数
func minDistance(word1 string, word2 string) int {
	// dp|i|j| 表示A的前i个字母和B的前j个字母之间的编辑距离
	// len(word1)==0, dp|0|j| = len(word2), 即要操作len(word2)次才能使word1==word2,
	// 操作中可删除AB的字母或可添加AB字母
	// dp|i|j| = min(dp|i-1|j|, dp|i|j-1|, dp|i-1|j-1|) + 1
	// 			or min(dp|i-1|j|, dp|i|j-1|, dp|i-1|j-1|-1) + 1 (word1[i] == word2[j] 如果相等就可以减少一次操作)
	// dp|i-1|j|, A的前i-1个和B的前j个，可在B后面加上Ai的字符，即dp|i|j|最小可以为dp|i-1|j|+1
	// dp|i|j-1|,
	// dp|i-1|j-1|
	// HORSE				ROS
	// A[1..i]				B[1..j]
	// dp|1|1| = min(dp|0|1|, dp|1|0|, dp|0|0|) + 1 = min(1, 1, 0) + 1 = 1 需要操作一次
	n, m := len(word1), len(word2)
	if n*m == 0 {
		return n+m
	}
	dp := make([][]int, n+1)
	// 定义边界dp|0|0|=0
	for i := 0; i < n + 1; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}
	for i := 0; i < m + 1; i++ {
		dp[0][i] = i
	}
	for i:=1; i<n+1; i++ {
		for j:=1; j<m+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]-1)) + 1
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
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
