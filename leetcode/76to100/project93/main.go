package main

import (
	"fmt"
	"strconv"
)

//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
func restoreIpAddresses(s string) []string {
	// 用小数点正确分割
	res := []string{}
	totalNum := 4
	ipAddrs := make([]int, totalNum)

	var fn func(int, int)
	fn = func(total, start int) {
		if total == totalNum { // 遍历到最后一段
			if start == len(s) { // 遍历完字符串
				ipAddr := ""
				for i := 0; i < len(ipAddrs); i++ {
					ipAddr += strconv.Itoa(ipAddrs[i])
					if i < len(ipAddrs) - 1 {
						ipAddr += "."
					}
				}
				res = append(res, ipAddr)
			}
			return
		}

		if start == len(s) {
			// 遍历完但是不够ip的长度
			return
		}

		if s[start] == '0' {
			// 该ip段以0开头的话，这一段只能为0
			ipAddrs[total] = 0
			fn(total+1, start+1)
		}

		addr := 0
		for end := start; end < len(s); end++ {
			addr = addr * 10 + int(s[end]-'0')
			if addr > 0 && addr <= 255 {
				ipAddrs[total] = addr
				fn(total+1, end+1)
			} else {
				break
			}
		}
	}
	fn(0, 0)
	return res
}



func main() {
	fmt.Println(restoreIpAddresses("010010"))
}
