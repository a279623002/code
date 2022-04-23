package main

import (
	"fmt"
	"math"
)

//输入：s = "42"
//输出：42
//解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。
//第 1 步："42"（当前没有读入字符，因为没有前导空格）
//^
//第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//^
//第 3 步："42"（读入 "42"）
//^
//解析得到整数 42 。
//由于 "42" 在范围 [-231, 231 - 1] 内，最终结果为 42 。
func myAtoi(s string) int {
	automation := NewAutomation()
	for i := 0 ; i < len(s) ; i++ {
		automation.get(s[i])
	}
	return automation.ans * automation.sign
}

type Automation struct {
	ans   int
	sign  int
	state string
	rule  map[string][]string
}

func (a *Automation) get(c uint8) {
	// 从当前状态根据字符 c 转移到下一个状态
	// 如开始状态 c 为空格，下一状态为开始
	// 数字状态 c为空格 下一状态结束
	a.state = a.rule[a.state][getCol(c)]

 	if a.state == "number" {
		a.ans = a.ans * 10 + int(c - '0')
		if a.sign == 1 {
			a.ans = min(a.ans, math.MaxInt32)
		} else {
			a.ans = min(a.ans, -math.MinInt32)
		}
	} else if a.state == "sign" {
		if c == '-' {
			a.sign = -1
		} else {
			a.sign = 1
		}
	}
}

func NewAutomation() *Automation {
	return &Automation{
		ans:   0,
		sign:  1,
		state: "start",
		rule: map[string][]string{
			"start":  {"start", "sign", "number", "end"},
			"sign":   {"end", "end", "number", "end"},
			"number": {"end", "end", "number", "end"},
			"end":    {"end", "end", "end", "end"},
		},
	}
}

func getCol(c uint8) (res int) {
	if c == ' ' {
		res = 0
	} else if c == '-' || c == '+' {
		res = 1
	} else if c >= '0' && c <= '9' {
		res = 2
	} else {
		res = 3
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(myAtoi(" -2333abc123")) // -2333
}
