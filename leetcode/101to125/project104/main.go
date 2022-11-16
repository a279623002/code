package project100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉树，找出其最大深度。
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
// 说明: 叶子节点是指没有子节点的节点。
// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，
// 返回它的最大深度 3 。
func maxDepth(root *TreeNode) int {
	return fn(root, 0)
}

func fn(root *TreeNode, n int) int {
	if root == nil {
		return n
	}

	return max(fn(root.Right, n+1), fn(root.Left, n+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
