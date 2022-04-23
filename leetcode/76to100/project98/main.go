package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	// 左子树节点右节点要大于当前节点并要小于根节点
	// 右子树节点左节点要小于当前节点并要大于根节点
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func main() {
	t := &TreeNode{Val:5}
	t.Left = &TreeNode{Val:4}
	t.Right = &TreeNode{Val:6}
	t.Right.Left = &TreeNode{Val:3}
	t.Right.Right = &TreeNode{Val:7}
	fmt.Println(isValidBST(t))
}
