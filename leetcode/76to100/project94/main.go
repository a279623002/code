package main

import "fmt"

// 		1
// 	   2 3
// 45 		67
// [4 2 5 1 6 3 7]
func inorderTraversal(root *TreeNode) (res []int) {
	var fn func (*TreeNode)
	fn = func(tree *TreeNode) {
		if tree == nil {
			return
		}
		fn(tree.Left)
		res = append(res, tree.Val)
		fn(tree.Right)
	}
	fn(root)
	return res
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {

	tree := &TreeNode{Val:1}
	tree.Left = &TreeNode{Val:2}
	tree.Right = &TreeNode{Val:3}
	tree.Left.Left = &TreeNode{Val:4}
	tree.Left.Right = &TreeNode{Val:5}
	tree.Right.Left = &TreeNode{Val:6}
	tree.Right.Right = &TreeNode{Val:7}
	res := inorderTraversal(tree)
	fmt.Println(res)
}