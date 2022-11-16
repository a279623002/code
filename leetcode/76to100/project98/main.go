package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
// 有效 二叉搜索树定义如下：
//
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树
// 输入：root = [2,1,3]
// 输出：true
// 输入：root = [5,1,4,null,null,3,6]
// 输出：false
// 解释：根节点的值是 5 ，但是右子节点的值是 4 。
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
	t := &TreeNode{Val: 5}
	t.Left = &TreeNode{Val: 4}
	t.Right = &TreeNode{Val: 6}
	t.Right.Left = &TreeNode{Val: 3}
	t.Right.Right = &TreeNode{Val: 7}
	fmt.Println(isValidBST(t))
}
