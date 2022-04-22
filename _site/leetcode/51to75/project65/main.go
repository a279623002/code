package main

import "fmt"

// 有效数字如下
// ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]
// 无效数字如下
//["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]
func isNumber(s string) bool {
	n := len(s)
	if s[0] == 'e' || s[0] == 'E' || s[n-1] == '-' || s[n-1] == '+' {
		return false
	}
	a := NewAutomation()
	for i := 0; i < n; i++ {
		a.state = a.rule[a.state][a.GetCol(s[i])]
	}
	if n == 1 {
		return a.state == "number"
	}
	// 1. true .1. false
	return a.state == "number" || (a.hasNum && a.state == ".") || a.state == "enum"
}

type Automation struct {
	state string
	enum bool
	xnum bool
	hasNum bool
	rule  map[string][]string
}

func NewAutomation() *Automation {
	return &Automation{
		state: "start",
		enum: false,
		xnum: false,
		hasNum: false,
		rule: map[string][]string{
			"start":  {"number", ".", "sign", "e", "enum", "end", "end", "end"},
			"number": {"number", ".", "end", "e", "enum", "end", "end", "end"},
			".":      {"number", "end", "end", "e","end", "end", "xnum", "end"},
			"sign":   {"number", ".", "end", "end","end", "end", "end", "end"},
			"e":      {"end", "end", "esign", "end", "enum", "end", "end", "end"},
			"enum":   {"end", "end", "end", "end", "enum", "end", "end", "end"},
			"esign":   {"end", "end", "end", "end", "enum", "end", "end", "end"},
			"xnum":   {"end", "end", "end", "e", "enum", "end", "xnum", "end"},
			"end":    {"end", "end", "end", "end", "end", "end", "end", "end"},
		},
	}
}

func (a *Automation) GetCol(c uint8) int {
	if c >= '0' && c <= '9' {
		if a.enum {
			if !a.hasNum {
				return 7
			}
			return 4
		}
		a.hasNum = true
		return 0
	} else if c == '.' {
		if a.state == "." {
			return 7
		}
		if a.xnum {
			return 6
		}
		a.xnum = true
		return 1
	} else if c == '-' || c == '+' {
		if a.state == "e" {
			if a.enum {
				return 4
			}
			return 7
		}
		return 2
	} else if c == 'e' || c == 'E' {
		a.enum = true
		return 3
	} else {
		return 7
	}
}

func main() {
	fmt.Println(isNumber("092e359-2"))
}
