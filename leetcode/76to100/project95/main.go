package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 二叉树 整个树每个节点满足左小右大
//输入：n = 3
//输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTree := helper(start, i-1) // 获取小于i的所有左树
		rightTree := helper(i+1, end) // 获取大于i的所有右树
		for _, left := range leftTree {
			for _, right := range rightTree {
				curTree := &TreeNode{Val:i} // 构建以i为头节点的二叉树
				curTree.Left = left
				curTree.Right = right
				res = append(res, curTree)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(generateTrees(1))
}
