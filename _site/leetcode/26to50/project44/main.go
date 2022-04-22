package main

import "fmt"

//输入:
//s = "aa"
//p = "*"
//输出: true
//解释: '*' 可以匹配任意字符串。
func isMatch(s string, p string) bool {
	// f|i|j| 表示s 的前i个字符与p 的前j个字符匹配的结果
	// if p[j-1] != '*'
	// 		f|i|j| = f|i-1|j-1| or false	s[0] = a p[0] = a
	// else
	// 		f|i|j| = f|i-1|j| or f|i|j-1| 	s[1] = b p[1] = * //i = 2, j = 2
	//										使用星号 		f|i-1|j|
	//										不使用星号 	f|i|j-1|
	//	f[0][0]=True，即当字符串 ss 和模式 pp 均为空时，匹配成功；
	//	f[i][0]=False，即空模式无法匹配非空字符串；
	//  f[0][j] 需要分情况讨论：因为星号才能匹配空字符串，所以只有当模式 pp 的前 j 个字符均为星号时，f[0][j] 才为真。

	sl, pl := len(s), len(p)
	f := make([][]bool, sl+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, pl + 1)
	}
	f[0][0] = true

	// 如果p的第一个字符是*,则设为true
	for i := 1; i <= pl; i++ {
		if p[i-1] == '*' {
			f[0][i] = true
		} else {
			break
		}
	}

	for i := 1; i <= sl; i++ {
		for j := 1; j <= pl; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j-1] || f[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				f[i][j] = f[i-1][j-1]
			}
		}
	}

	return f[sl][pl]
}

func main() {
	fmt.Println(isMatch("ba", "*a"))
}
