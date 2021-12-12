package main

import "fmt"

//输入：s = "()[]{}"
//输出：true
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	dict := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	res := []byte{}
	for i := 0; i < len(s); i++ {
		if v, ok := dict[s[i]]; ok {
			if len(res) == 0 || res[len(res) - 1] != v {
				return false
			}
			res = res[0:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}
	return len(res) == 0
}

func main() {
	fmt.Println(isValid("[{()}]"))
}
