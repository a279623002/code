package main

import (
	"fmt"
	"sort"
)

//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
func groupAnagrams(strs []string) [][]string {
	// sort.slice 快速排序，作为map的键，填入值
	dict := map[string][]string{}
	for _, v := range strs {
		str := []byte(v)
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		dict[string(str)] = append(dict[string(str)], v)
	}

	res := [][]string{}
	for _, v := range dict {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "ran"}))
}
