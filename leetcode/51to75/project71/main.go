package main

import (
	"fmt"
	"strings"
)

//输入：path = "/a/./b/../../c/"
//输出："/c"
func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	res := []string{}
	for _, v := range strs {
		if v == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else if v != "." && len(v) > 0 {
			res = append(res, v)
		}
	}
	return "/"+strings.Join(res, "/")
}

func main() {
	fmt.Println(simplifyPath("../a/./b/..//../c/"))
}
