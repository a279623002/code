package main

import (
	"fmt"
	"math"
)

//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
func minWindow(s string, t string) string {
	tmap, cmap := make(map[byte]int), make(map[byte]int)
	for k, _ := range t {
		tmap[t[k]]++
	}
	check := func() bool {
		for k, v := range tmap {
			if cmap[k] < v {
				return false
			}
		}
		return true
	}

	slow, fast, l, r, search := 0, 0, -1, -1, math.MaxInt32
	for fast < len(s) {
		if tmap[s[fast]] > 0 {
			cmap[s[fast]]++
		}
		for check() && slow <= fast {
			// 更新最小区间值
			if fast-slow+1 < search {
				search = fast - slow + 1
				l = slow
				r = slow + search // fast+1
			}
			// slow更新到第二个匹配值
			if _, ok := cmap[s[slow]]; ok {
				cmap[s[slow]]--
			}
			slow++
		}
		fast++
	}
	if l == -1 {
		return ""
	}
	return s[l:r]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
