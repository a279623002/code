package main

import "fmt"

//输入：s = "aa" p = "a*"
//输出：true
//解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

func isMatch(s string, p string) bool {
	// f|i|j| 表示s 的前i个字符与p 的前j个字符匹配的结果
	// if p[j-1] != '*'
	// 		f|i|j| = f|i-1|j-1| or false	match(i, j) s[i-1] p[i-1]
	// else
	// 		f|i|j| = f|i-1|j| or f|i|j-2| 	match(i, j-1) s[i-1] p[i-2]
	// 			s[i-1]=a, p[j-2]=? (p[j-1]=*)
	// 			f|i|j-2|为匹配0次的结果, s[i-2]=a,s[i-1]=a,	p[j-3]=a,p[j-2]=b, p[j-1]=*, 即s[i-1], p[j-3]的结果
	match := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' || s[i-1] == p[j-1] {
			return true
		}
		return false
	}
	f := make([][]bool, len(s)+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, len(p)+1)
	}
	f[0][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j - 1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if match(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if match(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[len(s)][len(p)]
}

func main() {
	fmt.Println(isMatch("a", "*a"))
}
