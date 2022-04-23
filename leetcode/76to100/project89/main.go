package main

import "fmt"

//输入：n = 2
//输出：[0,1,3,2]
//输入：n = 4
//输出：[0,1,3,2,6,7,5,4,12,13,15,14,10,11,9,8]
//				   i    ri  j   i     carry
//				=> 7 => 0   4   3	  	4
//				=> 6 => 1   3   3		4
//				=> 5 => 2   2   3		4
//				=> 4 => 3   1   3		4

//				=> 3 => 0	2	2		2
//				=> 2 => 1	1   2		2

//				=> 1 => 0	1   1		1

//				=> 0 => 0   0	0
func grayCode(n int) []int {
	res := make([]int, 1<<n)
	for i := 0; i < n; i++ {
		carry := 1 << i
		for j := 1; j < carry+1; j++ {
			//j+carry-1  => 1, 2, 3, 4, 5, 6, 7...
			//carry-j	 => 0, 1, 0, 3, 2, 1, 0...
			res[j+carry-1] = res[carry-j] + carry
		}
	}
	return res
}

func main() {
	fmt.Println(grayCode(4))
}
