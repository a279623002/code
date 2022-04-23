package main

import "fmt"

//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {
	res := []string{}
	back(&res, "", n, n)
	return res
}

// n=3, left=3, right=3
// 																left=2, right=3 (
// 							left=1, right=3 (( 													left=2, right=2 ()
// 			left=0, right=3 ((( 				left=1, right=2 (() 							left=1, right=2 ()(
// 			left=0, right=2 ((() 	left=0,right=2 (()( 		left=1, right=1 (()) 			left=1, right=1 ()()
// 			left=0, right=1 ((()) 	left=0,right=1 (()() 		left=0, right=1 (())( 			left=0, right=1 ()()(
// 			left=0, right=0 ((())) 	left=0,right=0 (()()) 		left=0, right=0 (())() 			left=0, right=0 ()()()
func back(res *[]string, cur string, left, right int) {
	//回溯跳出条件，
	//并不需要判断左括号是否用完，因为右括号生成的条件 right > left ，
	//所以右括号用完了就意味着左括号必定用完了
	if right == 0 {
		*res = append(*res, cur)
		return
	}

	if left > 0 {
		back(res, cur+"(", left-1, right)
	}

	// 括号成对存在，有左括号才会有右括号
	if right > left {
		back(res, cur+")", left, right-1)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
