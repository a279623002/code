package main

import "fmt"

//输入：n = 4, k = 2
//输出：
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]
func combine(n int, k int) [][]int {
	res := [][]int{}
	carry := []int{}
	var fn func(int)
	fn = func(cur int) {
		if cur > n+1 {
			return
		}
		// 剪枝
		if len(carry) == k {
			res = append(res, append([]int(nil), carry...))
			return
		}
		// 正常分枝伸展
		//if cur == n+1 {
		//	res = append(res, append([]int(nil), carry...))
		//	return
		//}
		carry = append(carry, cur)
		fn(cur + 1)
		carry = carry[:len(carry)-1]
		fn(cur + 1)
	}
	fn(1)
	return res
}

func combine1(n int, k int) [][]int {
	arr := []int{}
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	var res [][]int
	var fn func(int, []int, int)
	// n是k    i是数组[i:]
	fn = func(n int, carry []int, i int) {
		if n == 0 {
			res = append(res, append([]int{}, carry...))
		} else {
			if i < len(arr) {
				for key, v := range arr[i:] {
					carry = append(carry, v)
					// key+i+1 跳过已append 的下标
					// key每次递归一开始是从0算起，需要加上i
					// 1 2 3 4 当前为1 下一次递归为 2 3 4，
					fn(n-1, carry, key+i+1)
					carry = carry[:len(carry)-1]
				}
			}
		}
	}
	fn(k, []int{}, 0)
	return res
}

func main() {
	fmt.Println(combine(4, 2))
}
