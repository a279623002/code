package main

import "fmt"

func fn(nums []int) [][]int {
	res := [][]int{}

	carry := []int{}
	var f func(int)
	f = func(n int) {
		if n == len(nums) {
			res = append(res, append([]int(nil), carry...))
			return
		}
		carry = append(carry, nums[n])
		f(n+1)
		carry = carry[:len(carry)-1]
		f(n+1)
	}
	f(0)
	return res
}

func fn1(nums []int) [][]int {
	res := [][]int{}

	carry := []int{}
	var f func(int, bool)
	f = func(n int, ok bool) {
		if n == len(nums) {
			res = append(res, append([]int(nil), carry...))
			return
		}
		if !ok && n > 0 && nums[n-1] == nums[n] {
			return
		}

		carry = append(carry, nums[n])
		f(n+1, true)
		carry = carry[:len(carry)-1]
		f(n+1, false)
	}
	f(0, false)
	return res
}

func main() {
	fmt.Println(fn([]int{1, 2, 3, 4}))
	//fmt.Println(fn1([]int{1, 2, 2}))
}
